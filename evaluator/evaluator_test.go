package evaluator

import (
	"github.com/nel215/monkey/lexer"
	"github.com/nel215/monkey/object"
	"github.com/nel215/monkey/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if evaluated != tt.expected {
			t.Fatalf("fail")
		}
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}
