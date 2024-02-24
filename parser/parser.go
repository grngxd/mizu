package parser

import (
	"fmt"
	"mizu/lexer"
)

type Parser struct {
	l    *lexer.Lexer
	cur  lexer.Token
	peek lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	// Lexer should already have been initialized with tokens by the time it gets here
	// so we can just return a new parser with the lexer
	p := &Parser{l: l}
	l.Pos = 0

	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := NewProgram()

	p.NextToken()

	for p.cur.Type != lexer.EOF {
		var s Statement

		switch p.cur.Type {
		case lexer.CONST:
			s = p.parseDeclarationStatement(program)
		}
		program.Statements = append(program.Statements, s)
		p.NextToken()
	}
	return program
}

func (p *Parser) parseExpression(program *Program) Expression {
	switch p.cur.Type {
	case lexer.NUMBER:
		return p.parseNumber(program)
	case lexer.IDENTIFIER:
		return p.parseIdentifier(program)
	}
	return nil
}

func (p *Parser) parseNumber(program *Program) Expression {
	return &Number{Token: p.cur}
}

func (p *Parser) parseIdentifier(program *Program) Expression {
	return &Identifier{Token: p.cur, Value: p.cur.Literal}
}

func (p *Parser) parseDeclarationStatement(program *Program) *DeclarationStatement {
	_ = &DeclarationStatement{Token: p.cur}
	// We know the current token is a CONST token
	// it can either be const x int = 5 or const x := 5; so we need to check the next token
	// to see if it's an identifier or a type
	if p.peek.Type != lexer.IDENTIFIER {
		// We have an error
		program.Error(fmt.Errorf("expected an identifier after constant declaration"), p.cur.Line, p.cur.Col)
		return nil
	}

	p.NextToken()
	n := p.cur.Literal

	// We know the current token is an IDENTIFIER token
	// it can either be const x int = 5 or const x := 5 so we need to check the next token
	// to see if it's a type or a walrus
	switch p.peek.Type {
	case lexer.WALRUS:
		p.NextToken()
		p.NextToken()

		t := &Identifier{Token: lexer.Token{
			Type:    lexer.IDENTIFIER,
			Literal: "__INFER__",
			Line:    p.cur.Line,
			Col:     p.cur.Col,
		}, Value: "__INFER__"}

		// We know the current token is now the constant value
		v := p.parseExpression(program)
		if v == nil {
			// We have an error
			program.Error(fmt.Errorf("expected an expression after walrus in constant declaration"), p.cur.Line, p.cur.Col)
			return nil
		}

		return &DeclarationStatement{
			Token:      p.cur,
			Identifier: &Identifier{Token: lexer.Token{Type: lexer.IDENTIFIER, Literal: n, Line: p.cur.Line, Col: p.cur.Col}, Value: n},
			Value:      v,
			Type:       t,
		}
	case lexer.IDENTIFIER:
		p.NextToken()
		t := &Identifier{Token: p.cur, Value: p.cur.Literal}
		p.NextToken()
		p.NextToken()
		v := p.parseExpression(program)
		if v == nil {
			// We have an error
			program.Error(fmt.Errorf("expected an expression after type in constant declaration"), p.cur.Line, p.cur.Col)
			return nil
		}

		return &DeclarationStatement{
			Token:      p.cur,
			Identifier: &Identifier{Token: lexer.Token{Type: lexer.IDENTIFIER, Literal: n, Line: p.cur.Line, Col: p.cur.Col}, Value: n},
			Value:      v,
			Type:       t,
		}
	default:
		// We have an error
		program.Error(fmt.Errorf("expected a type or a walrus (:=) after identifier in constant declaration"), p.cur.Line, p.cur.Col)
	}

	return nil
}
