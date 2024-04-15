package token

import "fmt"

type literalValue map[string]int

const (

	// Keywords

	SELECT     = "SELECT"
	FROM       = "FROM"
	WHERE      = "WHERE"
	IN         = "IN"
	AND        = "AND"
	OR         = "OR"
	CREATE     = "CREATE"
	TABLE      = "TABLE"
	CHAR       = "CHAR"
	NUMERIC    = "NUMERIC"
	NOT        = "NOT"
	NULL       = "NULL"
	CONSTRAINT = "CONSTRAINT"
	KEY        = "KEY"
	PRIMARY    = "PRIMARY"
	FOREIGN    = "FOREIGN"
	REFERENCES = "REFERENCES"
	INSERT     = "INSERT"
	INTO       = "INTO"
	VALUES     = "VALUES"

	// Delimiters

	COMMA  = ","
	DOT    = "."
	LPAREN = "("
	RPAREN = ")"
	QUOTE  = "'"

	// Constants

	A = "a"
	D = "D"

	// Operators

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	// Relationals

	EQ   = "="
	LT   = "<"
	GT   = ">"
	LTEQ = "<="
	GTEQ = ">="

	// Token Types

	KEYWORD = "KEYWORD"

	DELIMITER  = "DELIMITER"
	RELATIONAL = "RELATIONAL"
	OPERATOR   = "OPERATOR"

	CONSTANT   = "RELATIONAL"
	IDENTIFIER = "IDENTIFIER"

	CN = "NUMERIC CONSTANT"
	CA = "ALPHANUMERIC CONSTANT"

	EOQ = "$"
)

type TokenType string

type Token struct {
	Type   TokenType // SELECT
	Token  string    // "SELECT"
	Symbol string    // "s"
	Value  int       // 10
	Code   int       // 1
	Line   int       // Linea 1
}

// Toma como parametro el tipo de token, el token propiamente, el simbolo, el valor, el codigo y el numero de linea
func NewToken(ty TokenType, tk string, sy string, va int, co int, ln int) *Token {
	return &Token{Type: ty, Token: tk, Symbol: sy, Value: va, Code: co, Line: ln}
}
func (tk *Token) String() string {
	return fmt.Sprintf("Type:%v\nToken:%v\nSymbol:%v\nValue:%v\nCode:%v\nLine:%v\n------\n", tk.Type, tk.Token, tk.Symbol, tk.Value, tk.Code, tk.Line)
}

// func LookupIdent(ident string) TokenType {
// 	if keyword, ok := keywords[ident]; ok {
// 		return keyword
// 	}
// 	return IDENT
// }
