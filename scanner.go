package main

import (
	"errors"
	"strconv"
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
}

func (scanner *Scanner) ScanTokens() []Token {
	var tokens []Token
	for !scanner.isAtEnd() {
		scanner.start = scanner.current
		token, err := scanner.scanToken()
		if err == nil {
			tokens = append(tokens, token)
		}
	}
	tokens = append(tokens, Token{tokenType: EOF, lexeme: "", literal: "", line: scanner.line})
	return tokens
}

func (scanner *Scanner) scanToken() (Token, error) {
	c := scanner.advance()
	switch c {
	case '(':
		return scanner.addToken(LEFT_PAREN, ""), nil
	case ')':
		return scanner.addToken(RIGHT_PAREN, ""), nil
	case '{':
		return scanner.addToken(LEFT_BRACE, ""), nil
	case '}':
		return scanner.addToken(RIGHT_BRACE, ""), nil
	case ',':
		return scanner.addToken(COMMA, ""), nil
	case '.':
		return scanner.addToken(DOT, ""), nil
	case '-':
		return scanner.addToken(MINUS, ""), nil
	case '+':
		return scanner.addToken(PLUS, ""), nil
	case ';':
		return scanner.addToken(SEMICOLON, ""), nil
	case '*':
		return scanner.addToken(STAR, ""), nil
	case '!':
		if scanner.match('=') {
			return scanner.addToken(BANG_EQUAL, ""), nil
		}
		return scanner.addToken(BANG, ""), nil
	case '=':
		if scanner.match('=') {
			return scanner.addToken(EQUAL_EQUAL, ""), nil
		}
		return scanner.addToken(EQUAL, ""), nil
	case '<':
		if scanner.match('=') {
			return scanner.addToken(LESS_EQUAL, ""), nil
		}
		return scanner.addToken(LESS, ""), nil
	case '>':
		if scanner.match('=') {
			return scanner.addToken(GREATER_EQUAL, ""), nil
		}
		return scanner.addToken(GREATER, ""), nil
	case '/':
		if scanner.match('/') {
			for scanner.peek() != '\n' && !scanner.isAtEnd() {
				scanner.advance()
			}
		} else {
			return scanner.addToken(SLASH, ""), nil
		}
	case ' ':
	case '\t':
	case '\r':
	case '\n':
		scanner.line++
	case '"':
		s := scanner.getString()
		return scanner.addToken(STRING, s), nil
	default:
		if isDigit(c) {
			num := scanner.getNumber()
			return scanner.addToken(NUMBER, num), nil
		} else if isAlpha(c) {
			scanner.getIdentifier()
			tokenType, ok := keywords[scanner.source[scanner.start:scanner.current]]
			if ok {
				return scanner.addToken(tokenType, ""), nil
			} else {
				return scanner.addToken(IDENTIFIER, ""), nil
			}
		} else {
			errorReport(uint(scanner.line), "Unexpected token")
		}
	}
	return Token{}, errors.New("not a token")
}

func (scanner *Scanner) moveCurrent() {
	scanner.current++
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) advance() byte {
	defer scanner.moveCurrent()
	return scanner.source[scanner.current]
}

func (scanner *Scanner) addToken(tokenType TokenType, literal any) Token {
	return Token{tokenType: tokenType, lexeme: scanner.source[scanner.start:scanner.current], line: scanner.line, literal: literal}
}

func (scanner *Scanner) match(expected rune) bool {
	if scanner.isAtEnd() || scanner.source[scanner.current] != byte(expected) {
		return false
	}
	return true
}

func (scanner *Scanner) peek() byte {
	if scanner.isAtEnd() {
		return byte(0)
	}
	return scanner.source[scanner.current]
}

func (scanner *Scanner) getString() string {
	for !scanner.isAtEnd() && scanner.peek() != '"' {
		if scanner.peek() == '\n' {
			scanner.line++
		}
		scanner.advance()
	}
	if scanner.isAtEnd() {
		errorReport(uint(scanner.line), "Unterminated string.")
	}
	scanner.advance()
	return scanner.source[scanner.start+1 : scanner.current-1]
}

func isDigit(c byte) bool {
	return c >= byte('0') && c <= byte('9')
}

func (scanner *Scanner) getNumber() float64 {
	for !scanner.isAtEnd() && isDigit(scanner.peek()) {
		scanner.advance()
	}
	if scanner.peek() == '.' && scanner.current+1 < len(scanner.source) && isDigit(scanner.source[scanner.current+1]) {
		scanner.advance()
		for !scanner.isAtEnd() && isDigit(scanner.peek()) {
			scanner.advance()
		}
	}
	num, _ := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 32)
	return num
}

func isAlpha(c byte) bool {
	return (c >= byte('a') && c <= byte('z')) ||
		(c >= byte('A') && c <= byte('Z')) ||
		(c == byte('_'))
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (scanner *Scanner) getIdentifier() {
	for isAlphaNumeric(scanner.peek()) {
		scanner.advance()
	}
}

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}
