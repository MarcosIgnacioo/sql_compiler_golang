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
			return token.SELECT, "s", 10, 1
		case "FROM":
			return token.FROM, "f", 11, 1
		case "WHERE":
			return token.WHERE, "w", 12, 1
		case "IN":
			return token.IN, "n", 13, 1
		case "AND":
			return token.AND, "y", 14, 1
		case "OR":
			return token.OR, "o", 15, 1
		case "CREATE":
			return token.CREATE, "c", 16, 1
		case "TABLE":
			return token.TABLE, "t", 17, 1
		case "CHAR":
			return token.CHAR, "h", 18, 1
		case "NUMERIC":
			return token.NUMERIC, "u", 19, 1
		case "NOT":
			return token.NOT, "e", 20, 1
		case "NULL":
			return token.NULL, "g", 21, 1
		case "CONSTRAINT":
			return token.CONSTRAINT, "b", 22, 1
		case "KEY":
			return token.KEY, "k", 23, 1
		case "PRIMARY":
			return token.PRIMARY, "p", 24, 1
		case "FOREIGN":
			return token.FOREIGN, "j", 25, 1
		case "REFERENCES":
			return token.REFERENCES, "l", 26, 1
		case "INSERT":
			return token.INSERT, "m", 27, 1
		case "INTO":
			return token.INTO, "q", 28, 1
		case "VALUES":
			return token.VALUES, "v", 29, 1
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
			return token.ASTERISK, "*", 72, 72
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
			// En caso de que el token no sea ninguna palabra o caracter reservado significa que puede ser 3 cosas
			// Una constante
			// Un identificador
			// Una constante numerica sin ''
			// Un token ilegal
			// Por lo que evaluamos las siguientes situaciones con regex
			// isConstant sirve para saber si es una constante
			isConstant := utils.StringMatchesAll(input, globals.REGEX_SQL_CONSTANT)
			if isConstant {
				constantCount++
				constantValue, constantType := getConstantValue(input)
				return constantType, input, constantValue, constantCount
			} else {
				// isIdentifier sirve para saber si es un identificador
				isIdentifier := utils.StringMatchesAll(input, globals.REGEX_SQL_IDENTIFIER)
				if isIdentifier {
					identifierCount++
					return IDENTIFIER, input, 4, identifierCount
				} else {
					// isDecimal sirve para saber si es un numero puramente, sin ''
					isDecimal := utils.StringMatchesAll(input, globals.REGEX_SQL_DECIMAL)
					if isDecimal {
						constantCount++
						return CONSTANT, input, 61, constantCount
					}
					// En el caso de que no encaje con ningun caso significa que es ilegal
					return ILLEGAL, input, 999, 999
				}
			}
		}
	}
}

// Definir si es una constante alfanumerica o numerica
func getConstantValue(input string) (int, token.TokenType) {
	isNumeric := utils.StringMatchesAll(input, globals.REGEX_SQL_NUMERICAL)
	if isNumeric {
		return 61, CN
	}
	return 62, CA
}
