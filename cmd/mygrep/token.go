package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	CHAR    = "CHAR"
	INT     = "INT"
	EOF     = ""

	BACKSLASH    = "\\"
	FORWARDSLASH = "/"
	DIGIT_EXPR   = "\\d"
	WORD_EXPR    = "\\w"
	DOLLAR_EXPR  = "$"
)
