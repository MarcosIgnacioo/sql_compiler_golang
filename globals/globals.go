package globals

const (
	// Regex para separar todas las palabras o caracteres
	REGEX_SQL_WORDS_SEPARATOR = `(‘|’|'|") ?[a-zA-Z@!\$\%#\\\(\)_\-0-9\[\]\{\}? ]*(’|'|‘|")|>=|<=|[a-zA-Z0-9\\$@#!\\-\\*]+|(,|.)`
	// Regex para saber si es una constante con ''
	REGEX_SQL_CONSTANT = `(‘|’|'|").*(’|'|‘|")$|\\d*`
	// Regex para saber si es una constante numerica
	REGEX_SQL_NUMERICAL = `(‘|’|'|")\d+(’|'|‘|")`
	// Regex para saber si es un identificador
	REGEX_SQL_IDENTIFIER = `[a-zA-Z#]*`
	// Regex para saber si es decimal
	REGEX_SQL_DECIMAL = `\d+`
)
