package parser

import (
	"fmt"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/stack"
	"github.com/MarcosIgnacioo/token"
)

func Parse(tokens arraylist.ArrayList) string {
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

				break
			}
		} else {
			productions := syntaxTable[currentRule][token.TokenType(currentToken.Token)]
			if productions != nil {
				if productions[0] != 99 {
					InsertInReverse(&rules, productions)
				}
			} else {
				fmt.Println("error")
				fmt.Println(currentRule)
				fmt.Println(productions)
				fmt.Println(currentToken.Type)
				fmt.Println(currentToken.String())
				break
			}
		}
	}
	return ""
}

func InsertInReverse(stack *stack.Stack, array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		stack.Push(array[i])
	}
}

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
		token.FROM: []int{99}, token.COMMA: []int{50, 302}, token.EOQ: []int{99},
	},

	304: {
		token.IDENTIFIER: []int{305},
	},

	305: {
		token.RELATIONAL: []int{99},
		token.FROM:       []int{99},
		token.IN:         []int{99},
		token.AND:        []int{99},
		token.OR:         []int{99},
		token.DELIMITER:  []int{99},
		token.RPAREN:     []int{99},
		token.EOQ:        []int{99},
	},

	306: {
		token.IDENTIFIER: []int{308, 307},
	},

	307: {
		token.WHERE:     []int{99},
		token.DELIMITER: []int{308, 307},
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
