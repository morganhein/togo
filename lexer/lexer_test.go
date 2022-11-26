package lexer

import (
	"testing"

	"togo/token"
)

func TestNextToken(t *testing.T) {
	input := `package main
import fmt
func main() {
		str := "Hello World!"
		fmt.Println("%v", str)
}
`
	tests := []struct {
		expectedType    token.Toke
		expectedLiteral string
	}{
		{token.PACKAGE, "package"},
		{token.IDENT, "main"},
		{token.IMPORT, "import"},
		{token.IDENT, "fmt"},
		{token.FUNC, "func"},
		{token.IDENT, "main"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.IDENT, "str"},
		{token.DEFINE, ":="},
		{token.IDENT, "\"Hello World!\""},
		{token.IDENT, "fmt"},
		{token.PERIOD, "."},
		{token.IDENT, "Println"},
		{token.LPAREN, "("},
		{token.IDENT, "\"%v\""},
		{token.IDENT, "str"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
