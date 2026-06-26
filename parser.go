package main

import (
	"slices"
	"strconv"
)

type Parser struct {
	Tokens []Token
}

func (Parser *Parser) get(index int) *Token {
	if len(Parser.Tokens) <= index {
		return &Parser.Tokens[len(Parser.Tokens)-1]
	}

	return &Parser.Tokens[index]
}

func (Parser *Parser) eat() Token {
	token := Parser.Tokens[0]

	if len(Parser.Tokens) <= 1 {
		return token
	}

	slices.Delete(Parser.Tokens, 0, 1)
	return token
}

func (Parser *Parser) parse() ASTNode {

}

func (Parser *Parser) parse_primary() ASTNode {
	token := Parser.eat()
	switch token.TokenType {
	case INT_TOKEN:
		val, err := strconv.ParseInt(token.Value, 0, 64)
		if err != nil {
			return NilNode{}
		}

		return IntNode{
			Value: int64(val),
		}
	case REAL_TOKEN:
		val, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return NilNode{}
		}

		return RealNode{
			Value: val,
		}
	}

	return NilNode{}
}
