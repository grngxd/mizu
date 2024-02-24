package main

import (
	"fmt"
	"mizu/lexer"
	"mizu/parser"
)

func main() {
	l := lexer.New(`const x int64 = 314`)
	l.Lex()
	for _, token := range l.Tokens {
		fmt.Println(token.String())
	}

	fmt.Println("-----------------------------------")

	p := parser.New(l)
	program := p.ParseProgram()
	for _, err := range program.Errors {
		fmt.Println(err)
	}

	fmt.Println(program.Literal())
}
