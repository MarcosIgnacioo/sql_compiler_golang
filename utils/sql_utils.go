package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarcosIgnacioo/arraylist"
	"github.com/MarcosIgnacioo/token"
)

const escape = "\x1b"

var ErrorQueries = []string{
	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,CARRERAS
WHERE ALUMNOS.A#=INSCRITOS.A# AND ALUMNOS.C#=CARRERAS.C#
AND INSCRITOS.SEMESTRE='2010I'
AND CARRERAS.CNOMBRE='ISC'
AND ALUMNOS.GENERACION='2010'`,
	`SELECT *
FROM PROFESORES
WHERE EDAD >45 AND GRADO=MAE' OR GRADO='DOC'`,
	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,
WHERE ALUMNOS.A#=INSCRITOS.A# AND
INSCRITOS.SEMESTRE='2010I'`,
	`SELECT ANOMBRE
FROM ALUMNOS
WHERE A# IN(SELECT A#
FROM INSCRITOS
WHERE P# IN (SELECT P#
FROM PROFESORES
WHERE GRADO='MAE'))

AND C# IN (SELECT C#
FROM
WHERE CNOMBRE='ISC')`,

	`SELECT ANOMBRE
FROM ALUMNOS A,INSCRITOS I,CARRERAS C
WHERE A.A#=I.A# AND A.C#=C.C#
AND I.SEMESTRE='2010I' C.CNOMBRE='ITC'`,
	`SELECT A#,ANOMBRE
FROM ALUMNOS
WHERE C# IN (SELECT C#
FROM CARRERAS
WHERE SEMESTRES=9)

AND A# (SELECT A#

FROM INSCRITOS
WHERE SEMESTRE='2010I')`,
	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,CARRERAS
WHERE ALUMNOS.A#=INSCRITOS.A# AND ALUMNOS.C#=CARRERAS.C#
AND INSCRITOS.SEMESTRE='2010I'
AND CARRERAS.CNOMBRE='ISC
AND ALUMNOS.GENERACION='2010'`,
	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,CARRERAS
WHERE ALUMNOS.A#=INSCRITOS.A# AND ALUMNOS.C#=CARRERAS.C#
AND INSCRITOS.SEMESTRE '2010I'
AND CARRERAS.CNOMBRE='ISC'
AND ALUMNOS.GENERACION='2010'`,
}

const (
	NONE = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
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

func Obtain_keys_from_syntax_table(hashmap map[token.TokenType][]int) []token.TokenType {
	keys := make([]token.TokenType, len(hashmap))
	i := 0
	for k := range hashmap {
		keys[i] = k
		i++
	}
	return keys
}

func ToInterfaceArray(array []token.TokenType) []interface{} {
	var interfaceArray []interface{}
	for _, item := range array {
		interfaceArray = append(interfaceArray, item)
	}
	return interfaceArray
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

func color(c int) string {
	if c == NONE {
		return fmt.Sprintf("%s[%dm", escape, c)
	}

	return fmt.Sprintf("%s[3%dm", escape, c)
}

func Format(c int, text string) string {
	return color(c) + text + color(NONE)
}
func PrintError(error string) {
	fmt.Println(Format(RED, error))
}

func PrintSuccess(error string) {
	fmt.Println(Format(GREEN, error))
}

func PrintQuery(query string, queryNumber int) {
	queryLines := strings.Split(query, "\n")
	queryHeader := fmt.Sprint("QUERY NO: ", strconv.Itoa(queryNumber))
	fmt.Println(Format(YELLOW, queryHeader))
	for i, ql := range queryLines {
		fmt.Println(i+1, ":", Format(BLUE, ql))
	}
}
