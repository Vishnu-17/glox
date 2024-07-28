package main

import "fmt"

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func (t TokenType) String() string {
	return [...]string{"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE",
		"COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "SLASH", "STAR", "BANG",
		"BANG_EQUAL", "EQUAL", "EQUAL_EQUAL", "GREATER", "GREATER_EQUAL",
		"LESS", "LESS_EQUAL", "IDENTIFIER", "STRING", "NUMBER", "AND",
		"CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "NIL", "OR",
		"PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE",
		"EOF"}[t]
}

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func (token *Token) ToString() {
	fmt.Printf("type: %s, lexeme: %s, lieral: %s, line: %d\n", token.tokenType.String(), token.lexeme, token.literal, token.line)
}
