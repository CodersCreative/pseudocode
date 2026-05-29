package main

type ASTNode interface {
	isNode()
}

type IFNode struct {
	Condition ASTNode
	Body      ASTNode
	Else      ASTNode
}

func (n IFNode) isNode() {}

type NilNode struct{}

func (n NilNode) isNode() {}
