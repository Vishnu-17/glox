package main

import (
	"errors"
	"strconv"

	"github.com/vishnu/glox/token"
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
}

func (scanner *Scanner) ScanTokens() []token.Token {
	var tokens []token.Token
	for !scanner.isAtEnd() {
		scanner.start = scanner.current
		token, err := scanner.scanToken()
		if err == nil {
			tokens = append(tokens, token)
		}
	}
	tokens = append(tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: "", Line: scanner.line})
	return tokens
}

func (scanner *Scanner) scanToken() (token.Token, error) {
	c := scanner.advance()
	switch c {
	case '(':
		return scanner.addToken(token.LEFT_PAREN, ""), nil
	case ')':
		return scanner.addToken(token.RIGHT_PAREN, ""), nil
	case '{':
		return scanner.addToken(token.LEFT_BRACE, ""), nil
	case '}':
		return scanner.addToken(token.RIGHT_BRACE, ""), nil
	case ',':
		return scanner.addToken(token.COMMA, ""), nil
	case '.':
		return scanner.addToken(token.DOT, ""), nil
	case '-':
		return scanner.addToken(token.MINUS, ""), nil
	case '+':
		return scanner.addToken(token.PLUS, ""), nil
	case ';':
		return scanner.addToken(token.SEMICOLON, ""), nil
	case '*':
		return scanner.addToken(token.STAR, ""), nil
	case '!':
		if scanner.match('=') {
			return scanner.addToken(token.BANG_EQUAL, ""), nil
		}
		return scanner.addToken(token.BANG, ""), nil
	case '=':
		if scanner.match('=') {
			return scanner.addToken(token.EQUAL_EQUAL, ""), nil
		}
		return scanner.addToken(token.EQUAL, ""), nil
	case '<':
		if scanner.match('=') {
			return scanner.addToken(token.LESS_EQUAL, ""), nil
		}
		return scanner.addToken(token.LESS, ""), nil
	case '>':
		if scanner.match('=') {
			return scanner.addToken(token.GREATER_EQUAL, ""), nil
		}
		return scanner.addToken(token.GREATER, ""), nil
	case '/':
		if scanner.match('/') {
			for scanner.peek() != '\n' && !scanner.isAtEnd() {
				scanner.advance()
			}
		} else {
			return scanner.addToken(token.SLASH, ""), nil
		}
	case ' ':
	case '\t':
	case '\r':
	case '\n':
		scanner.line++
	case '"':
		s := scanner.getString()
		return scanner.addToken(token.STRING, s), nil
	default:
		if isDigit(c) {
			num := scanner.getNumber()
			return scanner.addToken(token.NUMBER, num), nil
		} else if isAlpha(c) {
			scanner.getIdentifier()
			tokenType, ok := keywords[scanner.source[scanner.start:scanner.current]]
			if ok {
				return scanner.addToken(tokenType, ""), nil
			} else {
				return scanner.addToken(token.IDENTIFIER, ""), nil
			}
		} else {
			errorReport(uint(scanner.line), "Unexpected token")
		}
	}
	return token.Token{}, errors.New("not a token")
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

func (scanner *Scanner) addToken(tokenType token.TokenType, literal any) token.Token {
	return token.Token{Type: tokenType, Lexeme: scanner.source[scanner.start:scanner.current], Line: scanner.line, Literal: literal}
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

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}
