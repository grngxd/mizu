package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `const x := 5`
	l := New(input)
	l.Lex()

	expectedTokens := []Token{
		{Type: CONST, Literal: "const"},
		{Type: IDENTIFIER, Literal: "x"},
		{Type: WALRUS, Literal: ":="},
		{Type: NUMBER, Literal: "5"},
		{Type: EOF, Literal: ""},
	}

	for i, token := range l.Tokens {
		fmt.Println(token.String())
		if token.Type != expectedTokens[i].Type {
			t.Fatalf("Token[%d] - type wrong. expected=%q, got=%q", i, expectedTokens[i].Type, token.Type)
		}
		if token.Literal != expectedTokens[i].Literal {
			t.Fatalf("Token[%d] - literal wrong. expected=%q, got=%q", i, expectedTokens[i].Literal, token.Literal)
		}
	}
}
