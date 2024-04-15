package lexer

import (
	"regexp"
	"strings"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/globals"
	"github.com/MarcosIgnacioo/token"
	"github.com/MarcosIgnacioo/utils"
)

const (
	KEYWORD = "KEYWORD"

	DELIMITER  = "DELIMITER"
	RELATIONAL = "RELATIONAL"
	OPERATOR   = "OPERATOR"

	CONSTANT   = "CONSTANT"
	CN         = "NUMERIC CONSTANT"
	CA         = "ALPHANUMERIC CONSTANT"
	IDENTIFIER = "IDENTIFIER"

	ILLEGAL = "ILLEGAL"
	DOT     = "."
	RPAREN  = ")"
)

// Acepta un string (el query del usuario), lo separa por saltos de linea, y obtiene todos los tokens del query, los va clasificando con la funcion de LexQuery
func GetLexedTokens(query string) arraylist.ArrayList {
	r, _ := regexp.Compile(globals.REGEX_SQL_WORDS_SEPARATOR)
	queryLines := strings.Split(query, "\n")
	LexQuery := createLexerFunc()
	var tokens arraylist.ArrayList
	for i, v := range queryLines {
		words := r.FindAll([]byte(v), -1)
		wordsArray := utils.Bm_to_string_array(words)
		tokens = LexQuery(wordsArray, i+1)
	}
	lastToken := token.NewToken(token.EOQ, "$", "$", 199, 199, len(queryLines))
	tokens.Enqueue(lastToken)
	return tokens
}

// Genera la funcion para poder hacer el analisis lexico. La razon por la que se hace de esta manera es que asi mantenemos a nuestros tokens en un mismo arreglo durante todo. La funcion interna que produce simplemente lo que hace es aceptar un arreglo de palabras y un numero que sera la linea en la que se encuentre en ese momento el analizador lexico; Luego hara uso de la funcion proporcionada por generateLexerInstance para clasificar la palabra en un token, y una vez obtenidos los datos correspondientes al token creara un token con dichos datos.
func createLexerFunc() func([]string, int) arraylist.ArrayList {
	tokens := arraylist.NewArrayList(10)
	GetAttributes := generateLexerInstance()
	return func(words []string, ln int) arraylist.ArrayList {
		for _, word := range words {
			ty, sy, va, co := GetAttributes(word)
			tk := token.NewToken(ty, word, sy, va, co, ln)
			tokens.Enqueue(tk)
		}
		return tokens
	}
}

// Genera la funcion que sirve para clasificar a los tokens. Se hace de esta manera para poder tener un registro de los tokens en sus contadores propios. La funcion que se encarga de clasificar los tokens lo hace por medio de un switch statement, el cual toma el string que se le pase a la funcion y dependiendo de lo que tenga de valor lo clasificara por alguno de los tipos que pueda ser. Si no encaja con ningun tipo establecido significa que es ilegal
func generateLexerInstance() func(string) (token.TokenType, string, int, int) {
	identifierCount := 399
	constantCount := 599
	return func(input string) (token.TokenType, string, int, int) {
		switch input {
		case "SELECT":
			return KEYWORD, "s", 10, 1
		case "FROM":
			return KEYWORD, "f", 11, 1
		case "WHERE":
			return KEYWORD, "w", 12, 1
		case "IN":
			return KEYWORD, "n", 13, 1
		case "AND":
			return KEYWORD, "y", 14, 1
		case "OR":
			return KEYWORD, "o", 15, 1
		case "CREATE":
			return KEYWORD, "c", 16, 1
		case "TABLE":
			return KEYWORD, "t", 17, 1
		case "CHAR":
			return KEYWORD, "h", 18, 1
		case "NUMERIC":
			return KEYWORD, "u", 19, 1
		case "NOT":
			return KEYWORD, "e", 20, 1
		case "NULL":
			return KEYWORD, "g", 21, 1
		case "CONSTRAINT":
			return KEYWORD, "b", 22, 1
		case "KEY":
			return KEYWORD, "k", 23, 1
		case "PRIMARY":
			return KEYWORD, "p", 24, 1
		case "FOREIGN":
			return KEYWORD, "j", 25, 1
		case "REFERENCES":
			return KEYWORD, "l", 26, 1
		case "INSERT":
			return KEYWORD, "m", 27, 1
		case "INTO":
			return KEYWORD, "q", 28, 1
		case "VALUES":
			return KEYWORD, "v", 29, 1
		case ",":
			return DELIMITER, ",", 50, 5
		case ".":
			return DOT, ".", 51, 5
		case "(":
			return DELIMITER, "(", 52, 5
		case ")":
			return RPAREN, ")", 53, 5
		case "'":
			return DELIMITER, "'", 53, 5
		case "‘":
			return DELIMITER, "‘", 53, 5
		case "’":
			return DELIMITER, "’", 53, 5
		case "\"":
			return DELIMITER, "\"", 53, 5
		case "+":
			return OPERATOR, "+", 7, 70
		case "-":
			return OPERATOR, "-", 7, 71
		case "*":
			return OPERATOR, "*", 7, 72
		case "/":
			return OPERATOR, "/", 7, 73
		case "=":
			return RELATIONAL, "=", 8, 81
		case "<":
			return RELATIONAL, "<", 8, 82
		case ">":
			return RELATIONAL, ">", 8, 83
		case "<=":
			return RELATIONAL, "<=", 8, 84
		case ">=":
			return RELATIONAL, ">=", 8, 85
		default:
			isConstant := utils.StringMatchesAll(input, globals.REGEX_SQL_CONSTANT)
			if isConstant {
				constantCount++
				constantValue, constantType := getConstantValue(input)
				return constantType, input, constantValue, constantCount
			} else {
				isIdentifier := utils.StringMatchesAll(input, globals.REGEX_SQL_IDENTIFIER)
				if isIdentifier {
					identifierCount++
					return IDENTIFIER, input, 4, identifierCount
				} else {
					isDecimal := utils.StringMatchesAll(input, globals.REGEX_SQL_DECIMAL)
					if isDecimal {
						constantCount++
						return CONSTANT, input, 61, constantCount
					}
					return ILLEGAL, input, 999, 999
				}
			}
		}
	}
}

func getConstantValue(input string) (int, token.TokenType) {
	isNumeric := utils.StringMatchesAll(input, globals.REGEX_SQL_NUMERICAL)
	if isNumeric {
		return 61, CN
	}
	return 62, CA
}
