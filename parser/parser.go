package parser

import (
	"fmt"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/error"
	"github.com/MarcosIgnacioo/stack"
	"github.com/MarcosIgnacioo/token"
	"github.com/MarcosIgnacioo/utils"
)

const (
	CREATE = 200
	SELECT = 300
)

// Funcion para realizar el parseo de los tokens, se recibe un arreglo de tokens y se ejecuta el algoritmo, en caso de encontrar un error devuelve una variable de tipo error con los datos del propio error, si no encuentra ningun error sintactico devuelve un string que dice "success"
func Parse(tokens arraylist.ArrayList, query string) (*error.Error, *string) {

	userTokens := tokens.ArrayList
	rules := stack.New()
	start := userTokens[0].(*token.Token)

	// Super solucion bien gamer

	rules.Push(199)

	if start.Token == token.SELECT {
		rules.Push(300)
	} else if start.Token == token.CREATE || start.Token == token.INSERT {
		rules.Push(201)
	} else {
		error := error.New(query, start.Token, start.Line, []interface{}{"SELECT", "CREATE"})
		return error, nil
	}

	pointer := 0

	var currentRule int

	for currentRule != 199 {
		currentRule = rules.Pop().(int)
		currentToken := userTokens[pointer].(*token.Token)

		if currentRule < 200 {
			if currentRule == currentToken.Value {
				pointer++
			} else {
				fmt.Println("wtf")
				fmt.Println(currentRule)
				nextToken := userTokens[pointer+1].(*token.Token)
				fmt.Println(nextToken.String())
				fmt.Println(nextToken)
				error := error.New(query, currentToken.Token, currentToken.Line, []interface{}{token.LexerFinder[currentRule]})
				return error, nil
			}
		} else {
			productions := syntaxTable[currentRule][currentToken.Type]
			if productions != nil {
				if productions[0] != 99 {
					InsertInReverse(&rules, productions)
				}
			} else {
				expectedRules := utils.ToInterfaceArray(utils.Obtain_keys_from_syntax_table(syntaxTable[currentRule]))
				error := error.New(query, currentToken.Token, currentToken.Line, expectedRules)
				return error, nil
			}
		}
	}

	success := "QUERY VALIDO"
	return nil, &success

}

// Se inserta en reversa al stack
func InsertInReverse(stack *stack.Stack, array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		stack.Push(array[i])
	}
}

// Tabla sintactica con las reglas, cada regla contiene propiamente otro hashmap el cual contiene de llave al token terminal que se encuentre en ese momento, y su valor es un arreglo de enteros que son las producciones.
var syntaxTable = map[int]map[token.TokenType][]int{

	200: {
		token.CREATE: []int{16, 17, 4, 52, 202, 53, 55, 201},
	},

	201: {
		token.CREATE: []int{200},
		token.INSERT: []int{211},
		token.EOQ:    []int{99},
	},

	202: {
		token.IDENTIFIER: []int{4, 203, 52, 61, 53, 204, 205},
	},

	203: {
		token.CHAR: []int{18}, token.NUMERIC: []int{19},
	},

	204: {
		token.NOT: []int{20, 21}, token.COMMA: []int{99},
	},

	205: {
		token.COMMA: []int{50, 206}, token.DELIMITER: []int{50, 206}, token.RPAREN: []int{99},
	},

	206: {
		token.IDENTIFIER: []int{202},
		token.CONSTRAINT: []int{207},
	},

	207: {
		token.CONSTRAINT: []int{22, 4, 208, 52, 4, 53, 209},
	},

	208: {
		token.PRIMARY: []int{24, 23},
		token.FOREIGN: []int{25, 23},
	},

	209: {
		token.REFERENCES: []int{26, 4, 52, 4, 53, 210},
		token.COMMA:      []int{50, 207},
		token.RPAREN:     []int{99},
	},

	210: {
		token.COMMA:  []int{50, 207},
		token.RPAREN: []int{99},
	},

	211: {
		token.INSERT:     []int{27, 28, 4, 29, 52, 212, 53, 55, 215},
		token.IDENTIFIER: []int{213, 312},
	},

	212: {
		token.QUOTE:    []int{213, 214},
		token.CONSTANT: []int{213, 214},
		// Esto puede romper todo o arreglarlo no hay punto medio
		token.CN: []int{213, 214},
		token.CA: []int{213, 214},
	},

	213: {
		token.QUOTE: []int{54, 62, 54},
		// Esto puede romper todo o arreglarlo no hay punto medio
		token.CA: []int{62},
		token.CN: []int{61},
	},

	214: {
		token.COMMA:  []int{50, 212},
		token.RPAREN: []int{99},
	},

	215: {
		token.CREATE: []int{200},
		token.INSERT: []int{211},
		token.EOQ:    []int{99},
	},

	300: {
		token.SELECT: []int{10, 301, 11, 306, 310},
	},

	301: {
		token.IDENTIFIER: []int{302}, token.ASTERISK: []int{72},
	},

	302: {
		token.IDENTIFIER: []int{304, 303},
	},

	303: {
		token.FROM: []int{99}, token.COMMA: []int{50, 302}, token.EOQ: []int{99}, token.DELIMITER: []int{50, 302},
	},

	304: {
		token.IDENTIFIER: []int{4, 305},
	},

	305: {
		token.RELATIONAL: []int{99},
		token.FROM:       []int{99},
		token.IN:         []int{99},
		token.DELIMITER:  []int{99},
		// Ando rompiendo todee
		token.COMMA: []int{99},
		// Ando rompiendo todee
		token.AND:    []int{99},
		token.OR:     []int{99},
		token.DOT:    []int{51, 4},
		token.RPAREN: []int{99},
		token.EOQ:    []int{99},
	},

	306: {
		token.IDENTIFIER: []int{308, 307},
	},

	307: {
		token.WHERE:     []int{99},
		token.DELIMITER: []int{50, 306},
		// Ando rompiendo todo
		token.COMMA: []int{50, 306},
		// Ando rompiendo todo
		token.RPAREN: []int{99},
		token.EOQ:    []int{99},
	},

	308: {
		token.IDENTIFIER: []int{4, 309},
	},

	309: {
		token.IDENTIFIER: []int{4},
		token.WHERE:      []int{99},
		token.DELIMITER:  []int{99},
		token.COMMA:      []int{99},
		token.RPAREN:     []int{99},
		token.EOQ:        []int{99},
	},

	310: {
		token.WHERE:  []int{12, 311},
		token.RPAREN: []int{99},
		token.EOQ:    []int{99},
	},

	311: {
		token.IDENTIFIER: []int{313, 312},
	},

	312: {
		token.AND:    []int{317, 311},
		token.OR:     []int{317, 311},
		token.RPAREN: []int{99},
		token.EOQ:    []int{99},
	},

	313: {
		token.IDENTIFIER: []int{304, 314},
	},

	314: {
		token.RELATIONAL: []int{315, 316},
		token.IN:         []int{13, 52, 300, 53},
	},

	315: {
		token.RELATIONAL: []int{8},
	},

	316: {
		token.IDENTIFIER: []int{304},
		token.QUOTE:      []int{54, 318, 54},
		token.CA:         []int{318},
		token.CONSTANT:   []int{319},
		token.CN:         []int{319},
	},

	317: {
		token.AND: []int{14},
		token.OR:  []int{15},
	},

	318: {
		token.CA: []int{62},
	},

	319: {
		token.CN:       []int{61},
		token.CONSTANT: []int{61},
	},
}
