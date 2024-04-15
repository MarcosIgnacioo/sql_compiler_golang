package globals

const (
	REGEX_SQL_WORDS_SEPARATOR = `(‘|’|'|") ?[a-zA-Z@!\$\%#\\\(\)_\-0-9\[\]\{\}? ]*(’|'|‘|")|>=|<=|[a-zA-Z0-9\\$@#!\\-\\*]+|(,|.)`
	REGEX_SQL_CONSTANT        = `(‘|’|'|").*(’|'|‘|")$|\\d*`
	REGEX_SQL_NUMERICAL       = `(‘|’|'|")\d+(’|'|‘|")`
	REGEX_SQL_IDENTIFIER      = `[a-zA-Z#]*`
	REGEX_SQL_DECIMAL         = `\d+`
)
