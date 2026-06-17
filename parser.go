package main

type DataType interface {
	isType()
}

type BoolType struct{}

func (t BoolType) isType() {}

type IntType struct{}

func (t IntType) isType() {}

type RealType struct{}

func (t RealType) isType() {}

type StringType struct{}

func (t StringType) isType() {}

type CharType struct{}

func (t CharType) isType() {}

type DateType struct{}

func (t DateType) isType() {}

type ArrayType struct {
	From     uint32
	To       uint32
	DataType DataType
}

func (t ArrayType) isType() {}

type UserDefinedType struct {
	Identifier string
}

func (t UserDefinedType) isType() {}

type ASTNode interface {
	isNode()
}

type DateNode struct {
	Day   int8
	Month int8
	Year  int32
}

func (n DateNode) isNode() {}

type IntNode struct {
	Value int64
}

func (n IntNode) isNode() {}

type RealNode struct {
	Value float64
}

func (n RealNode) isNode() {}

type IdentifierNode struct {
	Value string
}

func (n IdentifierNode) isNode() {}

type StringNode struct {
	Value string
}

func (n StringNode) isNode() {}

type CharNode struct {
	Value rune
}

func (n CharNode) isNode() {}

type IFNode struct {
	Condition ASTNode
	Body      ASTNode
	Else      ASTNode
}

func (n IFNode) isNode() {}

type NilNode struct{}

func (n NilNode) isNode() {}

type BlockNode struct {
	Nodes []ASTNode
}

func (n BlockNode) isNode() {}

type ForNode struct {
	Identifier string
	From       ASTNode
	To         ASTNode
	Step       ASTNode
	Body       ASTNode
}

func (n ForNode) isNode() {}

type WhileNode struct {
	Condition ASTNode
	Body      ASTNode
}

func (n WhileNode) isNode() {}

type RepeatNode struct {
	Condition ASTNode
	Body      ASTNode
}

func (n RepeatNode) isNode() {}

type Param struct {
	Identifier string
	DataType   DataType
	ByRef      bool
}

type ProcNode struct {
	Identifier string
	Parameters []Param
	Body       ASTNode
}

func (n ProcNode) isNode() {}

type FuncNode struct {
	Identifier string
	Parameters []Param
	ReturnType DataType
	Body       ASTNode
}

func (n FuncNode) isNode() {}

type CallNode struct {
	Func      ASTNode
	Arguments []ASTNode
}

func (n CallNode) isNode() {}
