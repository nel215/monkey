package parser

import (
	"github.com/nel215/monkey/ast"
	"github.com/nel215/monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
	}

	for i, _ := range tests {
		stmt := program.Statements[i]
		_, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Fatalf("s not *ast.LetStatement. got=%T", stmt)
		}
	}

}