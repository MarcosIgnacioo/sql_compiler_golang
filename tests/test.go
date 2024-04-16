package tests

import (
	"fmt"

	"github.com/MarcosIgnacioo/lexer"
	"github.com/MarcosIgnacioo/parser"
	"github.com/MarcosIgnacioo/utils"
)

func TestLexer() {

	var input int
	queries := utils.ErrorQueries
	for {
		fmt.Println("Ingrese el numero del query que desea evaluar [1-8]\nEscriba 0 para salir")
		fmt.Scanln(&input)
		if input == 0 {
			break
		}
		// input = 6
		if input-1 < len(queries) {
			query := queries[input-1]
			tokens := lexer.GetLexedTokens(query)
			err, good := parser.Parse(tokens, query)
			utils.PrintQuery(query, input)
			if err != nil {
				utils.PrintError(err.String())
			} else {
				utils.PrintSuccess(*good)
			}
		}
		// break
	}
}
