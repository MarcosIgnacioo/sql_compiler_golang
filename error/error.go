package error

import (
	"fmt"
)

type Error struct {
	Query    string
	From     interface{}
	Line     int
	Expected []interface{}
	Message  string
}

func New(q string, f interface{}, ln int, e []interface{}) *Error {
	return &Error{Query: q, From: f, Line: ln, Expected: e}
}

func (e *Error) String() string {
	e.Message += fmt.Sprint("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n")
	e.Message += fmt.Sprint("Error ocurred in line: ", e.Line, " nearby: `", e.From, "`\n")
	e.Message += fmt.Sprint("Expected some of this tokens\n")
	for _, tok := range e.Expected {
		e.Message += fmt.Sprintf("`%v`\n", tok)
	}
	return e.Message
}
