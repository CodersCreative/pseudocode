package main

// --- Data Types ---

type UserType interface {
	isUserType()
}

type DataType interface {
	isType()
}

type BoolType struct{}

func (t BoolType) isType()     {}
func (t BoolType) isUserType() {}

type IntType struct{}

func (t IntType) isType()     {}
func (t IntType) isUserType() {}

type RealType struct{}

func (t RealType) isType()     {}
func (t RealType) isUserType() {}

type StringType struct{}

func (t StringType) isType()     {}
func (t StringType) isUserType() {}

type CharType struct{}

func (t CharType) isType()     {}
func (t CharType) isUserType() {}

type DateType struct{}

func (t DateType) isType()     {}
func (t DateType) isUserType() {}

type ArrayType struct {
	From     uint32
	To       uint32
	DataType DataType
}

func (t ArrayType) isType()     {}
func (t ArrayType) isUserType() {}

type Array2DType struct {
	FromX    uint32
	ToX      uint32
	FromY    uint32
	ToY      uint32
	DataType DataType
}

func (t Array2DType) isType()     {}
func (t Array2DType) isUserType() {}

type PointerType struct {
	DataType DataType
}

func (t PointerType) isType()     {}
func (t PointerType) isUserType() {}

type UserDefinedType struct {
	Identifier string
}

func (t UserDefinedType) isType()     {}
func (t UserDefinedType) isUserType() {}

// --- AST Nodes ---

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

type DeclareNode struct {
	Identifiers []string
	DataType    DataType
	Value       ASTNode
}

func (n DeclareNode) isNode() {}

type AssignNode struct {
	Identifier string
	Value      ASTNode
}

func (n AssignNode) isNode() {}

type IFNode struct {
	Condition ASTNode
	Body      ASTNode
	Else      ASTNode
}

func (n IFNode) isNode() {}

type CaseArm struct {
	Value ASTNode
	// To will only be used for ranges and Value will become a from value
	To     ASTNode
	Action ASTNode
}

type CaseNode struct {
	Arms []CaseArm
	Else ASTNode
}

func (n CaseNode) isNode() {}

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

type ClassProperty struct {
	IsPublic   bool
	Identifier string
	DataType   DataType
}

type ClassMethod struct {
	IsPublic bool
	Function FuncNode
}

type ClassNode struct {
	Identifier string
	Inherits   []string
	Properties []ClassProperty
	Methods    []ClassMethod
}

type Param struct {
	Identifier string
	DataType   DataType
	ByRef      bool
}

// A Proc can just be a FuncNode with a ReturnType
type FuncNode struct {
	Identifier string
	Parameters []Param
	ReturnType DataType
	Body       ASTNode
}

func (n FuncNode) isNode() {}

// TODO Ensure that builtin calls like INPUT get routed here so that INPUT() is also valid
type CallNode struct {
	Func      ASTNode
	Arguments []ASTNode
}

func (n CallNode) isNode() {}

type NotNode struct {
	Value ASTNode
}

func (n NotNode) isNode() {}

// --- Binary Operations ---

const (
	GTOperator = iota
	LTOperator
	GEqOperator
	LEqOperator
	EqOperator
	NEqOperator
	AndOperator
	OrOperator
)

type BinaryNode struct {
	Left     ASTNode
	Right    ASTNode
	Operator uint8
}

func (n BinaryNode) isNode() {}

// --- User-Defined Types ---

type TypeNode struct {
	Identifier string
	Value      UserType
}

func (n TypeNode) isNode() {}

type KeyValType struct {
	Key   string
	Value DataType
}

type RecordType struct {
	Properties []KeyValType
}

func (t RecordType) isUserType() {}

type EnumType struct {
	Variants []string
}

func (t EnumType) isUserType() {}

type SetType struct {
	DataType DataType
}

func (t SetType) isUserType() {}
