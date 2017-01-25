package parser

import (
	"github.com/nel215/monkey/ast"
	"github.com/nel215/monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, str := range tests {
		stmt := program.Statements[i]
		letStmt, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Fatalf("s not *ast.LetStatement. got=%T", stmt)
		}

		if letStmt.Name.Value != str.expectedIdentifier {
			t.Fatalf("letStmt.Name.Value not '%s'. got=%s", str.expectedIdentifier, letStmt.Name.Value)
		}

		if letStmt.Name.TokenLiteral() != str.expectedIdentifier {
			t.Fatalf("s.Name not '%s'. got=%s", str.expectedIdentifier, letStmt.Name)
		}
	}

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
