// ast.go
package parser

import "fmt"

type Node interface {
	Literal() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
	Errors     []error
}

func (p *Program) Literal() string {
	if len(p.Statements) == 0 {
		return ""
	}
	s := ""
	for _, statement := range p.Statements {
		s += statement.Literal() + "\n"
	}

	return s
}

func NewProgram() *Program {
	return &Program{Statements: []Statement{}, Errors: []error{}}
}

func (p *Program) Error(err error, line int, col int) {
	nerr := fmt.Errorf("Error: %s. %d:%d", err.Error(), line, col)
	p.Errors = append(p.Errors, nerr)
}
