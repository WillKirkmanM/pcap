package main

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	ASSIGN    = "="
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	DOT       = "."
	COMMA     = ","

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	PROGRAM = "PROGRAM"
	BEGIN   = "BEGIN"
	END     = "END"
	LET     = "LET"
	WRITELN = "WRITELN"
	READLN  = "READLN"
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

var keywords = map[string]TokenType{
	"program": PROGRAM,
	"begin":   BEGIN,
	"end":     END,
	"let":     LET,
	"writeln": WRITELN,
	"readln":  READLN,
}

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
