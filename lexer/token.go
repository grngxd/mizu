package lexer

import "fmt"

type Type string
type Token struct {
	Type    Type
	Literal string
	Line    int
	Col     int
}

const (
	// Special tokens
	ILLEGAL Type = "ILLEGAL"
	EOF          = "EOF"
	EOL          = "EOL"
	COMMENT      = "COMMENT"
	// Literals/Types
	IDENTIFIER = "IDENTIFIER" // variable names, function names, etc.
	NUMBER     = "NUMBER"     // 123
	BOOLEAN    = "BOOLEAN"    // true | false
	RUNE       = "RUNE"       // 'a'
	FUNCTION   = "FUNCTION"   // fn
	// Control flow
	IF       = "IF"       // if
	ELSE     = "ELSE"     // else
	RETURN   = "RETURN"   // return
	FOR      = "FOR"      // for
	WHILE    = "WHILE"    // while
	BREAK    = "BREAK"    // break
	CONTINUE = "CONTINUE" // continue
	NULL     = "NULL"     // null
	// Delimiters
	LPAREN  = "LPAREN"  // (
	RPAREN  = "RPAREN"  // )
	LBRACE  = "LBRACE"  // {
	RBRACE  = "RBRACE"  // }
	LBRACK  = "LBRACK"  // [
	RBRACK  = "RBRACK"  // ]
	COMMA   = "COMMA"   // ,
	DOT     = "DOT"     // . || used for object access OR floating point
	COLON   = "COLON"   // :
	DCOLON  = "DCOLON"  // ::
	SEMIC   = "SEMIC"   // ;
	QMARK   = "QMARK"   // ?
	DOUBLEQ = "DOUBLEQ" // "
	SINGLEQ = "SINGLEQ" // '
	// Operators
	WALRUS = "WALRUS" // := | used for variable declaration (when type is implicit)
	ASSIGN = "ASSIGN" // = | used for variable declaration (when type is explicit)
	PLUS   = "PLUS"   // +
	MINUS  = "MINUS"  // -
	MUL    = "MUL"    // *
	DIV    = "DIV"    // /
	MOD    = "MOD"    // %
	POW    = "POW"    // ^
	// Comparison
	EQUAL            = "EQUAL"            // ==
	STRICT_EQUAL     = "STRICT_EQUAL"     // ===
	NOT_EQUAL        = "NOT_EQUAL"        // !=
	STRICT_NOT_EQUAL = "STRICT_NOT_EQUAL" // !==
	GREATER          = "GREATER"          // >
	GREATER_EQUAL    = "GREATER_EQUAL"    // >=
	LESS             = "LESS"             // <
	LESS_EQUAL       = "LESS_EQUAL"       // <=
	// Logical
	AND = "AND" // &&
	OR  = "OR"  // ||
	NOT = "NOT" // !
	// Mizu specific
	IMPORT  = "IMPORT"
	CONST   = "CONST"
	DECLARE = "DECLARE"
)

var Keywords = map[string]Type{
	"true":     BOOLEAN,
	"false":    BOOLEAN,
	"fn":       FUNCTION,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"for":      FOR,
	"while":    WHILE,
	"break":    BREAK,
	"continue": CONTINUE,
	"null":     NULL,
	"import":   IMPORT,
	"const":    CONST,
	"declare":  DECLARE,
}

func NewToken(t Type, lit string, line, col int) Token {
	return Token{Type: t, Literal: lit, Line: line, Col: col}
}

func (t *Token) String() string {
	s := fmt.Sprintf(
		`{
	Type: "%s",
	Literal: "%s",
	Line: %d,
	Col: %d
}`, t.Type, t.Literal, t.Line, t.Col)

	return s
}
