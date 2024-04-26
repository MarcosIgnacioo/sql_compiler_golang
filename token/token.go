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

	COMMA     = ","
	DOT       = "."
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	QUOTE     = "'"

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

	CONSTANT   = "CONSTANT"
	IDENTIFIER = "IDENTIFIER"

	CN = "NUMERIC CONSTANT"
	CA = "ALPHANUMERIC CONSTANT"

	EOQ = "$"
)

// El tipo de los tokens (SELECT, COMMA, EQUALS, etc)
type TokenType string

// Struct que tendran los tokens
type Token struct {
	Type   TokenType // SELECT
	Token  string    // "SELECT"
	Symbol string    // "s"
	Value  int       // 10
	Code   int       // 1
	Line   int       // 1
}

// Toma como parametro el tipo de token, el token propiamente, el simbolo, el valor, el codigo y el numero de linea
func NewToken(ty TokenType, tk string, sy string, va int, co int, ln int) *Token {
	return &Token{Type: ty, Token: tk, Symbol: sy, Value: va, Code: co, Line: ln}
}

func (tk *Token) String() string {
	return fmt.Sprintf("Type:%v\nToken:%v\nSymbol:%v\nValue:%v\nCode:%v\nLine:%v\n------\n", tk.Type, tk.Token, tk.Symbol, tk.Value, tk.Code, tk.Line)
}

var LexerFinder = map[int]string{
	4:   "IDENTIFIER",
	10:  "SELECT",
	11:  "FROM",
	12:  "WHERE",
	13:  "IN",
	14:  "AND",
	15:  "OR",
	16:  "CREATE",
	17:  "TABLE",
	18:  "CHAR",
	19:  "NUMERIC",
	20:  "NOT",
	21:  "NULL",
	22:  "CONSTRAINT",
	23:  "KEY",
	24:  "PRIMARY",
	25:  "FOREIGN",
	26:  "REFERENCES",
	27:  "INSERT",
	28:  "INTO",
	29:  "VALUES",
	50:  ",",
	51:  ".",
	52:  "(",
	53:  ")",
	54:  "'",
	55:  ";",
	61:  "ALFANUMERICA",
	62:  "ALFANUMERICA",
	70:  "+",
	71:  "-",
	72:  "*",
	73:  "/",
	81:  ">",
	82:  "<",
	83:  "=",
	84:  ">=",
	85:  "<=",
	199: "FINAL DEL QUERY",
}
