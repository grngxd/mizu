// statements.go
package parser

import (
	"fmt"
	"mizu/lexer"
)

type DeclarationStatement struct {
	Statement
	Token      lexer.Token
	Identifier *Identifier
	Value      Expression
	Type       *Identifier
}

func (ds *DeclarationStatement) statementNode() {}
func (ds *DeclarationStatement) Literal() string {
	return fmt.Sprintf("DeclarationStatement: %s %s %s", ds.Identifier.Literal(), ds.Type.Literal(), ds.Value.Literal())
}
