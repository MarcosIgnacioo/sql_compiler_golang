package parser

import (
	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/error"
	"github.com/MarcosIgnacioo/stack"
	"github.com/MarcosIgnacioo/token"
	"github.com/MarcosIgnacioo/utils"
)

// Funcion para realizar el parseo de los tokens, se recibe un arreglo de tokens y se ejecuta el algoritmo, en caso de encontrar un error devuelve una variable de tipo error con los datos del propio error, si no encuentra ningun error sintactico devuelve un string que dice "success"
func Parse(tokens arraylist.ArrayList, query string) (*error.Error, *string) {

	userTokens := tokens.ArrayList
	rules := stack.New()
	rules.Push(199)
	rules.Push(300)
	pointer := 0

	var currentRule int

	for currentRule != 199 {

		currentRule = rules.Pop().(int)
		currentToken := userTokens[pointer].(*token.Token)

		if currentRule < 300 {
			if currentRule == currentToken.Value {
				pointer++
			} else {
				expectedRules := utils.ToInterfaceArray(utils.Obtain_keys_from_syntax_table(syntaxTable[currentRule]))
				error := error.New(query, currentToken.Token, currentToken.Line, expectedRules)
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
		token.AND:        []int{99},
		token.OR:         []int{99},
		token.DOT:        []int{51, 4},
		token.RPAREN:     []int{99},
		token.EOQ:        []int{99},
	},

	306: {
		token.IDENTIFIER: []int{308, 307},
	},

	307: {
		token.WHERE:     []int{99},
		token.DELIMITER: []int{50, 306},
		token.RPAREN:    []int{99},
		token.EOQ:       []int{99},
	},

	308: {
		token.IDENTIFIER: []int{4, 309},
	},

	309: {
		token.IDENTIFIER: []int{4},
		token.WHERE:      []int{99},
		token.DELIMITER:  []int{99},
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
