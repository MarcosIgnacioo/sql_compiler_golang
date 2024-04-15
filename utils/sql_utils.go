package utils

import (
	"regexp"

	"github.com/MarcosIgnacioo/arraylist"
)

// Aqui iran los queries
// Aqui estaran las regex
// TODO pasar esto al lexer
// func GetSQLWords(query string) {
// 	r, _ := regexp.Compile(globals.REGEX_SQL_WORDS_SEPARATOR)
// 	queryLines := strings.Split(query, "\n")
// 	for i, v := range queryLines {
// 		words := r.FindAll([]byte(v), -1)
// 		wordsArray := bm_to_string_array(words)
// 		lexer.LexQuery(wordsArray, i+1)
// 	}
// }

func Bm_to_string_array(matrix [][]byte) []string {
	sArray := arraylist.NewArrayList(10)
	for _, bytes := range matrix {
		word := string(bytes)
		byte := []byte(word)[0]
		if word != " " && byte != 9 {
			sArray.Enqueue(word)
		}
	}
	return sArray.ConvertToStringArray()
}

// Funcion auxiliar que sirve para saber si el string entero encaja con la regex que se le pase
func StringMatchesAll(word string, regex string) bool {
	r, _ := regexp.Compile(regex)
	evaluated := r.FindAll([]byte(word), -1)
	if len(evaluated) == 0 {
		return false
	}
	return len(evaluated[0]) == len([]byte(word))
}
