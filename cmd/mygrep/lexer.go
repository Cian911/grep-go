package main

import "fmt"

// Tokenize a byte string.
type Lexer struct {
	input        string
	position     int  // Current position in input (current char)
	readPosition int  // Current reading position in input (after current char): i.e The _next_ position.
	ch           byte // Current char being inspected
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '\\':
		if l.peekChar() == 'd' {
			ch := l.ch
			literal := string(ch) + string(l.ch)
			l.ch = byte(l.readPosition)
			tok = Token{Type: DIGIT_EXPR, Literal: literal}
		} else if l.peekChar() == 'w' {
			ch := l.ch
			literal := string(ch) + string(l.ch)
			l.ch = byte(l.readPosition)
			tok = Token{Type: WORD_EXPR, Literal: literal}
		} else {
			tok = l.newToken(BACKSLASH, l.ch)
		}
	default:
		fmt.Println(l.ch)
	}

	// Advance until end of input
	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// We've reached the end of our input, reset to 0
		// 0 is ASCII for NULL
		l.ch = 0
	} else {
		// Othereise, set ch to next value in input
		l.ch = l.input[l.readPosition]
	}
	// Set position to next value
	l.position = l.readPosition
	l.readPosition += 1
}

// Peek at the next character after the current position
func (l *Lexer) peekChar() byte {
	if l.readPosition <= len(l.input) {
		return l.input[l.readPosition]
	}

	return 0
}

func (l *Lexer) newToken(exp string, ch byte) Token {
	return Token{
		Type:    TokenType(exp),
		Literal: string(l.ch),
	}
}

func (l *Lexer) isDigit() bool {
	return '0' <= l.ch && l.ch <= '9'
}

func (l *Lexer) isLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || 'A' <= l.ch && l.ch <= 'Z' || l.ch == '_'
}
