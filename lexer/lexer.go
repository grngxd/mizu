package lexer

import (
	"unicode"
)

type Lexer struct {
	Source string
	Pos    int
	Line   int
	Col    int
	Char   rune
	Tokens []Token
}

func New(source string) *Lexer {
	l := &Lexer{
		Source: source,
		Line:   1,
		Col:    1,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.Pos >= len(l.Source) {
		l.Char = 0
	} else {
		l.Char = rune(l.Source[l.Pos])
	}
	l.Pos++
	if l.Char == '\n' {
		l.Tokens = append(l.Tokens, Token{Type: EOL, Literal: "\n", Line: l.Line, Col: l.Col})
		l.Line++
		l.Col = 0
	}
	l.Col++
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.Char) {
		l.readChar()
	}
}

func (l *Lexer) Peek() rune {
	if l.Pos >= len(l.Source) {
		return 0
	}
	return rune(l.Source[l.Pos])
}

func (l *Lexer) readIdentifier() string {
	start := l.Pos - 1
	for (unicode.IsLetter(l.Char) || unicode.IsDigit(l.Char) || l.Char == '_') && l.Char != '.' {
		l.readChar()
	}
	return l.Source[start : l.Pos-1]
}

func (l *Lexer) readNumber() string {
	start := l.Pos - 1
	for unicode.IsDigit(l.Char) {
		l.readChar()
	}
	return l.Source[start : l.Pos-1]
}

func (l *Lexer) NextToken() Token {
	var token Token
	l.skipWhitespace()

	switch l.Char {
	default:
		if unicode.IsLetter(l.Char) {
			literal := l.readIdentifier()
			tokenType, ok := Keywords[literal]
			if !ok {
				tokenType = IDENTIFIER
			}

			token = Token{Type: tokenType, Literal: literal, Line: l.Line, Col: l.Col}
			return token
		} else if unicode.IsDigit(l.Char) {
			literal := l.readNumber()
			token = Token{Type: NUMBER, Literal: literal, Line: l.Line, Col: l.Col}
		} else {
			//token = Token{Type: ILLEGAL, Literal: string(l.Char), Line: l.Line, Col: l.Col}
			switch l.Char {
			default:
				token = Token{Type: ILLEGAL, Literal: string(l.Char), Line: l.Line, Col: l.Col}
			case '(':
				token = Token{Type: LPAREN, Literal: "(", Line: l.Line, Col: l.Col}
			case ')':
				token = Token{Type: RPAREN, Literal: ")", Line: l.Line, Col: l.Col}
			case '{':
				token = Token{Type: LBRACE, Literal: "{", Line: l.Line, Col: l.Col}
			case '}':
				token = Token{Type: RBRACE, Literal: "}", Line: l.Line, Col: l.Col}
			case '[':
				token = Token{Type: LBRACK, Literal: "[", Line: l.Line, Col: l.Col}
			case ']':
				token = Token{Type: RBRACK, Literal: "]", Line: l.Line, Col: l.Col}
			case ',':
				token = Token{Type: COMMA, Literal: ",", Line: l.Line, Col: l.Col}
			case '.':
				token = Token{Type: DOT, Literal: ".", Line: l.Line, Col: l.Col}
			case ':':
				if l.Peek() == '=' {
					l.readChar()
					token = Token{Type: WALRUS, Literal: ":=", Line: l.Line, Col: l.Col}
				} else if l.Peek() == ':' {
					l.readChar()
					token = Token{Type: DCOLON, Literal: "::", Line: l.Line, Col: l.Col}
				} else {
					token = Token{Type: COLON, Literal: ":", Line: l.Line, Col: l.Col}
				}
			case ';':
				token = Token{Type: SEMIC, Literal: ";", Line: l.Line, Col: l.Col}
			case '?':
				token = Token{Type: QMARK, Literal: "?", Line: l.Line, Col: l.Col}
			case '"':
				token = Token{Type: DOUBLEQ, Literal: "\"", Line: l.Line, Col: l.Col}
			case '\'':
				token = Token{Type: SINGLEQ, Literal: "'", Line: l.Line, Col: l.Col}
			case '+':
				token = Token{Type: PLUS, Literal: "+", Line: l.Line, Col: l.Col}
			case '-':
				token = Token{Type: MINUS, Literal: "-", Line: l.Line, Col: l.Col}
			case '*':
				token = Token{Type: MUL, Literal: "*", Line: l.Line, Col: l.Col}
			case '/':
				token = Token{Type: DIV, Literal: "/", Line: l.Line, Col: l.Col}
			case '%':
				token = Token{Type: MOD, Literal: "%", Line: l.Line, Col: l.Col}
			case '^':
				token = Token{Type: POW, Literal: "^", Line: l.Line, Col: l.Col}
			case '=':
				if l.Peek() == '=' {
					l.readChar()
					token = Token{Type: EQUAL, Literal: "==", Line: l.Line, Col: l.Col}
				} else {
					token = Token{Type: ASSIGN, Literal: "=", Line: l.Line, Col: l.Col}
				}
			case '>':
				if l.Peek() == '=' {
					l.readChar()
					token = Token{Type: GREATER_EQUAL, Literal: ">=", Line: l.Line, Col: l.Col}
				} else {
					token = Token{Type: GREATER, Literal: ">", Line: l.Line, Col: l.Col}
				}
			case '!':
				if l.Peek() == '=' {
					l.readChar()
					token = Token{Type: NOT_EQUAL, Literal: "!=", Line: l.Line, Col: l.Col}
				} else {
					token = Token{Type: ILLEGAL, Literal: "!", Line: l.Line, Col: l.Col}
				}
			case '<':
				if l.Peek() == '=' {
					l.readChar()
					token = Token{Type: LESS_EQUAL, Literal: "<=", Line: l.Line, Col: l.Col}
				} else {
					token = Token{Type: LESS, Literal: "<", Line: l.Line, Col: l.Col}
				}
			}
		}
	case 0:
		l.Tokens = append(l.Tokens, Token{Type: EOL, Literal: "\n", Line: l.Line, Col: l.Col})
		token = Token{Type: EOF, Literal: "", Line: l.Line, Col: l.Col}
	}

	l.readChar()
	return token
}

func (l *Lexer) Lex() {
	for {
		token := l.NextToken()
		l.Tokens = append(l.Tokens, token)
		if token.Type == EOF {
			break
		}
	}
}
