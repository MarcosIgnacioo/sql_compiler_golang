package tests

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		if input == 9 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Escriba su query")
			queryUser, error := reader.ReadString('\n')

			if error != nil {
				fmt.Println("Error leyendo", error)
				return
			}

			queryUser = strings.TrimSpace(queryUser)
			tokens := lexer.GetLexedTokens(queryUser)

			err, good := parser.Parse(tokens, queryUser)
			utils.PrintQuery(queryUser, input)
			if err != nil {
				utils.PrintError(err.String())
			} else {
				utils.PrintSuccess(*good)
			}
			continue
		}
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
