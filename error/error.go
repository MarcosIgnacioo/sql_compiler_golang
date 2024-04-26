package error

import (
	"fmt"
)

// Struct para los errores
type Error struct {
	Query    string
	From     interface{}
	Line     int
	Expected []interface{}
	Message  string
}

// Funcion para crear un error
func New(q string, f interface{}, ln int, e []interface{}) *Error {
	return &Error{Query: q, From: f, Line: ln, Expected: e}
}

// Funcnion para obtenter el texto del error
func (e *Error) String() string {
	e.Message += fmt.Sprint("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n")
	if e.From == "$" {
		e.From = "FINAL DEL QUERY"
	}
	e.Message += fmt.Sprint("Error ocurrio en la linea: ", e.Line, " cerca de: `", e.From, "`\n")
	e.Message += fmt.Sprint("Se esperaban algunos de los siguientes tokens\n")
	for _, tok := range e.Expected {
		e.Message += fmt.Sprintf("`%v`\n", tok)
	}
	return e.Message
}
