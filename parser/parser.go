package parser

import (
	"github.com/nel215/monkey/lexer"
	"github.com/nel215/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	return p
}
