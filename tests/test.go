package tests

import (
	"github.com/MarcosIgnacioo/lexer"
	"github.com/MarcosIgnacioo/parser"
)

func TestLexer() {
	query := `SELECT *
	FROM ALUMNOS,INSCRITOS,CARRERAS
	WHERE ALUMNOS.A#=INSCRITOS.A# AND ALUMNOS.C#=CARRERAS.C#
	AND INSCRITOS.SEMESTRE='2010I'
	AND CARRERAS.CNOMBRE='ISC'
	AND ALUMNOS.GENERACION=2010`
	tokens := lexer.GetLexedTokens(query)
	parser.Parse(tokens)
}
