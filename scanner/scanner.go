package scanner

import (
	"fmt"

	"github.com/zbsss/glox/errors.go"
	"github.com/zbsss/glox/tokens"
)

type (
	Scanner struct {
		source string
		tokens []*tokens.Token

		start   int // start location of current lexeme
		current int // current location in lexeme
		line    int
	}
)

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  make([]*tokens.Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []*tokens.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, &tokens.Token{Type: tokens.EOF, Line: s.line})

	return s.tokens
}

func (s *Scanner) scanToken() {
	char := s.advance()

	switch char {
	case '(':
		s.addToken(tokens.LEFT_PAREN, nil)
	case ')':
		s.addToken(tokens.RIGHT_PAREN, nil)
	case '{':
		s.addToken(tokens.LEFT_BRACE, nil)
	case '}':
		s.addToken(tokens.RIGHT_BRACE, nil)
	case ',':
		s.addToken(tokens.COMMA, nil)
	case '.':
		s.addToken(tokens.DOT, nil)
	case '-':
		s.addToken(tokens.MINUS, nil)
	case '+':
		s.addToken(tokens.PLUS, nil)
	case ';':
		s.addToken(tokens.SEMICOLON, nil)
	case '*':
		s.addToken(tokens.STAR, nil)
	case '!':
		s.addToken(tokens.BANG, nil)
	case '=':
		s.addToken(tokens.EQUAL, nil)
	case '<':
		s.addToken(tokens.LESS, nil)
	case '>':
		s.addToken(tokens.GREATER, nil)
	default:
		errors.Error(s.line, fmt.Sprintf("unexpected character: %c", char))
	}
}

func (s *Scanner) addToken(tokenType tokens.TokenType, literal interface{}) {
	s.tokens = append(s.tokens, &tokens.Token{
		Type:    tokenType,
		Lexeme:  s.source[s.start:s.current],
		Literal: literal,
		Line:    s.line,
	})
}

func (s *Scanner) advance() rune {
	s.current++
	return rune(s.source[s.current-1])
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
