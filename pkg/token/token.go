package token

// all token types
const (
	ILLEGAL = "ILLEGAL" // token/character we don't know about
	EOF = "EOF" // end of file, tells the parser it can stop

	// identifiers + literals
	IDENTIFIER = "IDENTIFIER" // add, x, y
	INTEGER = "INTEGER"

	// operators
	ASSIGN = "="
	PLUS = "+"

	// delimiters
	COMMA = ","
	SEMICOLON = ";"
	LEFT_PAREN = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE = "{"
	RIGHT_BRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}