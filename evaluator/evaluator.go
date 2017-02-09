package evaluator

import (
	"github.com/nel215/monkey/ast"
	"github.com/nel215/monkey/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		for _, statement := range node.Statements {
			return Eval(statement)
		}
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

		// Expression
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}
	return nil
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
