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
	// QUERY 1 El error es el simbolo [

	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,CARRERAS
WHERE ALUMNOS.A#=INSCRITOS.A# AND ALUMNOS.C#=CARRERAS.C#
AND INSCRITOS.SEMESTRE='2010I'
AND CARRERAS.CNOMBRE='ISC'
AND ALUMNOS.GENERACION='2010'`,

	// QUERY 2 El error es el ' seguido de MAE

	`SELECT *
FROM PROFESORES
WHERE EDAD >45 AND GRADO=MAE' OR GRADO='DOC'`,

	// QUERY 3 El error es la , despues de INSCRITOS

	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,NOMBRE
WHERE ALUMNOS.A#=INSCRITOS.A# AND
INSCRITOS.SEMESTRE='2010I'`,

	// QUERY 4 El error que no hay nada despues del FROM

	`SELECT ANOMBRE
FROM ALUMNOS
WHERE A# IN(SELECT A#
FROM INSCRITOS
WHERE P# IN (SELECT P#
FROM PROFESORES
WHERE GRADO='MAE'))

AND C# (SELECT C#
FROM
WHERE CNOMBRE='ISC')`,

	// QUERY 5 El error es que no tiene un and despues de '2010I'

	`SELECT ANOMBRE
FROM ALUMNOS A,INSCRITOS I,CARRERAS C
WHERE A.A#=I.A# AND A.C#=C.C#
AND I.SEMESTRE='2010I' OR C.CNOMBRE='ITC'`,

	// QUERY 6 El error es despues de AND A# se esperaba IN

	`SELECT A#,ANOMBRE
FROM ALUMNOS
WHERE C# IN (SELECT C#
FROM CARRERAS
WHERE SEMESTRES=9)

AND A# IN (SELECT A#

FROM INSCRITOS
WHERE SEMESTRE='2010I')`,

	// QUERY 7 El error es que falta un delimitador despues de 'ISC

	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,CARRERAS
WHERE ALUMNOS.A#=INSCRITOS.A# AND ALUMNOS.C#=CARRERAS.C#
AND INSCRITOS.SEMESTRE='2010I'
AND CARRERAS.CNOMBRE='ISC
AND ALUMNOS.GENERACION='2010'`,

	// QUERY 8 El error es que se esperaba un signo de = despuesde INSCRITOS.SEMESTRE y antes de '2010I'

	`SELECT ANOMBRE
FROM ALUMNOS,INSCRITOS,CARRERAS
WHERE ALUMNOS.A#=INSCRITOS.A# AND ALUMNOS.C#=CARRERAS.C#
AND INSCRITOS.SEMESTRE = '2010I'
AND CARRERAS.CNOMBRE='ISC'
AND ALUMNOS.GENERACION='2010'`,

	// QUERY 9 El error es que dice numerico en vez de numeric, no tiene coma antes de la constraint y no tiene punto y coma al final

	`CREATE TABLE DEPARTAMENTOS(
D# CHAR(2) NOT NULL,
DNOMBRE NUMERICO(6) NOT NULL
CONSTRAINT PK_DEPARTAMENTOS PRIMARY KEY (D#))`,

	// QUERY 10 El error es que le falta una coma en su tercera linea despues del not null, le falta un parentesis al final y un ;

	`CREATE TABLE DEPARTAMENTOS(
D# CHAR(2) NOT NULL,
DNOMBRE NUMERIC(6) NOT NULL
CONSTRAINT PK_DEPARTAMENTOS PRIMARY KEY (D#)`,

	// QUERY 11 El error es que dice numerico en vez de numeric, una coma despues del segundo not null no tiene identificador en el parentesis de hasta abajito y un ;

	`CREATE TABLE DEPARTAMENTOS(
D# CHAR(2) NOT NULL,
DNOMBRE NUMERIC(6) NOT NULL
CONSTRAINT PK_DEPARTAMENTOS PRIMARY KEY ())`,

	// QUERY 12 El error es que no tiene INTO en la linea 5
	`INSERT INTO INSCRITOS VALUES ('R01','A1','M1','P3','M','2010I',60);
INSERT INTO INSCRITOS VALUES ('R02','A1','M5','P4','M','2011I',75);
INSERT INTO INSCRITOS VALUES ('R03','A1','M2','P3','V','2010I',78);
INSERT INTO INSCRITOS VALUES ('R04','A1','M5','P4','M','2011II',80);
INSERT INSCRITOS VALUES ('R05','A2','M3','P6','V','2010I',86);
INSERT INTO INSCRITOS VALUES ('R06','A2','M4','P7','V','2010I',90);
INSERT INTO INSCRITOS VALUES ('R07','A3','M1','P2','M','2011I',70);
INSERT INTO INSCRITOS VALUES ('R08','A3','M5','P9','V','2011II',82);
`,

	// QUERY 13 El error es que antes del 'P2' no hay una coma
	`INSERT INTO INSCRITOS VALUES ('R01','A1','M1','P3','M','2010I',60);
INSERT INTO INSCRITOS VALUES ('R02','A1','M5','P4','M','2011I',75);
INSERT INTO INSCRITOS VALUES ('R03','A1','M2','P3','V','2010I',78);
INSERT INTO INSCRITOS VALUES ('R04','A1','M5','P4','M','2011II',80);
INSERT INTO INSCRITOS VALUES ('R05','A2','M3','P6','V','2010I',86);
INSERT INTO INSCRITOS VALUES ('R06','A2','M4','P7','V','2010I',90);
INSERT INTO INSCRITOS VALUES ('R07','A3','M1' 'P2','M','2011I',70);
INSERT INTO INSCRITOS VALUES ('R08','A3','M5','P9','V','2011II',82);`,

	// QUERY 14 El error es en la linea 3 hay una coma sobrante o se esperaba otra constante y en P2 de nuevo no hay una coma antes de el
	`INSERT INTO INSCRITOS VALUES ('R01','A1','M1','P3','M','2010I',60);
INSERT INTO INSCRITOS VALUES ('R02','A1','M5','P4','M','2011I',75);
INSERT INTO INSCRITOS VALUES ('R03','A1','M2','P3','V','2010I',);
INSERT INTO INSCRITOS VALUES ('R04','A1','M5','P4','M','2011II',80);
INSERT INTO INSCRITOS VALUES ('R05','A2','M3','P6','V','2010I',86);
INSERT INTO INSCRITOS VALUES ('R06','A2','M4','P7','V','2010I',90);
INSERT INTO INSCRITOS VALUES ('R07','A3','M1' 'P2','M','2011I',70);
INSERT INTO INSCRITOS VALUES ('R08','A3','M5','P9','V','2011II',82);`,

	// QUERY 15 El error esta en que no hay una coma antes depues del not null y antes del constraint

	`CREATE TABLE DEPARTAMENTOS(
D# CHAR(2) NOT NULL,
DNOMBRE CHAR(6) NOT NULL,
CONSTRAINT PK_DEPARTAMENTOS PRIMARY KEY (D#));

INSERT INTO DEPARTAMENTOS VALUES ('D1','CIECOM');
INSERT INTO DEPARTAMENTOS VALUES ('D2','CIETIE');
INSERT INTO DEPARTAMENTOS VALUES ('D3','CIEING');
INSERT INTO DEPARTAMENTOS VALUES ('D4','CIEECO');
INSERT INTO DEPARTAMENTOS VALUES ('D5','CIEBAS');

CREATE TABLE CARRERAS(
C# CHAR(2) NOT NULL,
CNOMBRE CHAR(3) NOT NULL,
VIGENCIA CHAR(4) NOT NULL,
SEMESTRES NUMERIC(2) NOT NULL,
D# CHAR(2) NOT NULL,
CONSTRAINT PK_CARRERAS PRIMARY KEY (C#),
CONSTRAINT FK_CARRERAS FOREIGN KEY (D#) REFERENCES DEPARTAMENTOS(D#));

INSERT INTO CARRERAS VALUES ('C1','ISC','2010',9,'D1');
INSERT INTO CARRERAS VALUES ('C2','LIN','2008',9,'D1');
INSERT INTO CARRERAS VALUES ('C3','ICI','2009',10,'D2');
INSERT INTO CARRERAS VALUES ('C4','IIN','2010',8,'D3');
INSERT INTO CARRERAS VALUES ('C5','LAE','2011',8,'D4');
INSERT INTO CARRERAS VALUES ('C6','ARQ','2010',10,'D2');

CREATE TABLE ALUMNOS(
A# CHAR(2) NOT NULL,
ANOMBRE CHAR(20) NOT NULL,
GENERACION CHAR(4) NOT NULL,
SEXO CHAR(1) NOT NULL,
C# CHAR(2) NOT NULL,
CONSTRAINT PK_ALUMNOS PRIMARY KEY(A#),
CONSTRAINT FK_ALUMNOS FOREIGN KEY(C#)REFERENCES CARRERAS(C#));

INSERT INTO ALUMNOS VALUES ('A1','ALBA JESSICA','2009','F','C1');
INSERT INTO ALUMNOS VALUES ('A2','CARREY JIM','2010','M','C2');
INSERT INTO ALUMNOS VALUES ('A3','JOLIE ANGELINE','2011','F','C1');
INSERT INTO ALUMNOS VALUES ('A4','SMITH WILL','2012','M','C6');
INSERT INTO ALUMNOS VALUES ('A5','MESSI LIONEL','2010','M','C1');
INSERT INTO ALUMNOS VALUES ('A6','ESPINOZA PAOLA','2012','F','C6');
INSERT INTO ALUMNOS VALUES ('A7','BLANCO CUAHTEMOC','2009','M','C5');
INSERT INTO ALUMNOS VALUES ('A8','SANCHEZ HUGO','2010','M','C4');

CREATE TABLE MATERIAS(
M# CHAR(2) NOT NULL,
MNOMBRE CHAR(6) NOT NULL,
CREDITOS NUMERIC(2) NOT NULL,
C# CHAR(2) NOT NULL,
CONSTRAINT PK_MATERIAS PRIMARY KEY (M#),
CONSTRAINT FK_MATERIAS FOREIGN KEY (C#) REFERENCES CARRERAS(C#));`,
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

// Convertir matriz de bytes a un arreglo de string
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

// Funcion para obtener las keys de un hashmap
func Obtain_keys_from_syntax_table(hashmap map[token.TokenType][]int) []token.TokenType {
	keys := make([]token.TokenType, len(hashmap))
	i := 0
	for k := range hashmap {
		keys[i] = k
		i++
	}
	return keys
}

// Funcion para convertir a un array de tipo generico
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

// Seccion de texto de color

func color(c int) string {
	if c == NONE {
		return fmt.Sprintf("%s[%dm", escape, c)
	}

	return fmt.Sprintf("%s[3%dm", escape, c)
}

func Format(c int, text string) string {
	return color(c) + text + color(NONE)
}

// Funcion para imprimir cuando hay un error

func PrintError(error string) {
	fmt.Println(Format(RED, error))
}

// Funcion para imprimir cuando no hay errores
func PrintSuccess(success string) {
	fmt.Println(Format(GREEN, success))
}

// Funcion para imprimir un query
func PrintQuery(query string, queryNumber int) {
	queryLines := strings.Split(query, "\n")
	queryHeader := fmt.Sprint("QUERY NO: ", strconv.Itoa(queryNumber))
	fmt.Println(Format(YELLOW, queryHeader))
	for i, ql := range queryLines {
		fmt.Println(i+1, ":", Format(BLUE, ql))
	}
}
