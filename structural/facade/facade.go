package main

type Scanner struct {
	inputStream string // Keep it simple
}

func (scanner Scanner) Scan() Token {
}

type Token struct {
}

type Parser struct {
}

func (parser Parser) parse(scanner Scanner, ProgramNodeBuilder ProgramNodeBuilder) {

}

type ProgramNodeBuilder struct {
	node ProgramNode
}

func (programNodeBuilder ProgramNodeBuilder) newVariable(variableName string) ProgramNode {
}

func (programNodeBuilder ProgramNodeBuilder) newAssignment(variable ProgramNode, expression ProgramNode) ProgramNode {
}

func (programNodeBuilder ProgramNodeBuilder) newReturnStatement(value ProgramNode) ProgramNode {
}

func (programNodeBuilder ProgramNodeBuilder) newCondition(condition ProgramNode, truePart ProgramNode, falsePart ProgramNode) ProgramNode {
}

type ProgramNode interface {
	getSourcePosition(line int, index int)
	add(ProgramNode)
	remove(ProgramNode)
	traverse(CodeGenerator)
}

type ExpressionNode struct {
	children []ProgramNode
}

func (node ExpressionNode) getSourcePosition(line int, index int) {
}

func (node ExpressionNode) add(programNode ProgramNode) {	
}

func (node ExpressionNode) remove(programNode ProgramNode) {
}

func (node ExpressionNode) traverse(codeGenerator CodeGenerator) {
	codeGenerator.visit(node)

	for i := 0; i < len(node.children); i++ {
		node.children[i].traverse(codeGenerator)
	}
}

type CodeGenerator struct {
	output string
} 

func (codeGenerator CodeGenerator) visit(programNode ProgramNode) {

}

type Compiler struct {
}

func (compiler Compiler) compile(input string, output string) {
	scanner := new(Scanner)
	programNodeBuilder := new(ProgramNodeBuilder)
	parser := new(Parser)

	parser.parse(*scanner, *programNodeBuilder)
	
	codeGenerator := CodeGenerator {
		output: output,
	}
	parseTree := programNodeBuilder.node
	parseTree.traverse(codeGenerator)
}