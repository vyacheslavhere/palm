package core

/*
 * Represents kind of the token
 */
type TokenKind string

const (
	Lparen TokenKind = "("
	Rparen TokenKind = ")"
	Lbrace TokenKind = "{"
	Rbrace TokenKind = "}"
	Let    TokenKind = "Let"
	Do     TokenKind = "Do"
	Echo   TokenKind = "Echo"
	Id     TokenKind = "Id"
	Run    TokenKind = "Run"
	String TokenKind = "String"
	Concat TokenKind = "<>"
)
