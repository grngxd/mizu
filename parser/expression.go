// expression.go - This file contains the expression parser for the mizu language.
package parser

import "mizu/lexer"

type Identifier struct {
	Token lexer.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Literal() string { return i.Token.Literal }

type Number struct {
	Expression
	Token lexer.Token
	Value int
}

func (n *Number) expressionNode() {}
func (n *Number) Literal() string { return n.Token.Literal }
