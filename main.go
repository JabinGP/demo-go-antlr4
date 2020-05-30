package main

import (
	"log"
	"strconv"

	"github.com/JabinGP/demo-go-antlr4/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// EvalVisitor 自定义访问器，继承自BaseCalVisitor
type EvalVisitor struct {
	*parser.BaseCalVisitor
	memory map[string]int
}

// VisitAssign ...
func (e *EvalVisitor) VisitAssign(ctx parser.AssignContext) int {
	id := ctx.ID().GetText()
	value := e.Visit(ctx.Expr()).(int)
	e.memory[id] = value
	return value
}

// VisitPrintExpr ...
func (e *EvalVisitor) VisitPrintExpr(ctx parser.PrintExprContext) int {
	value := e.Visit(ctx.Expr()).(int)
	log.Println(value)
	return 0
}

// VisitInt ...
func (e *EvalVisitor) VisitInt(ctx parser.IntContext) int {
	text := ctx.INT().GetText()
	intValue, err := strconv.Atoi(text)
	if err != nil {
		log.Println(err)
		return 0
	}
	return intValue
}

// VisitId ...
func (e *EvalVisitor) VisitId(ctx parser.IdContext) int {
	id := ctx.ID().GetText()

	if value, ok := e.memory[id]; ok {
		return value
	}

	return 0
}

// VisitMulDiv ...
func (e *EvalVisitor) VisitMulDiv(ctx parser.MulDivContext) int {
	left := e.Visit(ctx.Expr(0)).(int)
	right := e.Visit(ctx.Expr(1)).(int)

	result := left * right

	switch ctx.GetOp().GetTokenType() {
	case parser.CalParserMUL:
		result = left * right
	case parser.CalParserSUB:
		result = left / right
	default:
		result = left / right
	}

	return result
}

// VisitAddSub ...
func (e *EvalVisitor) VisitAddSub(ctx parser.AddSubContext) int {
	left := e.Visit(ctx.Expr(0)).(int)
	right := e.Visit(ctx.Expr(1)).(int)

	var result int
	switch ctx.GetOp().GetTokenType() {
	case parser.CalParserADD:
		result = left + right
	case parser.CalParserSUB:
		result = left - right
	default:
		result = left - right
	}

	return result
}

// VisitParens ...
func (e *EvalVisitor) VisitParens(ctx parser.ParensContext) int {
	return e.Visit(ctx.Expr()).(int)
}
func main() {

	sourseFileName := "./input"

	input, err := antlr.NewFileStream(sourseFileName)
	if err != nil {
		panic(err)
	}
	lexer := parser.NewCalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewCalParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	tree := p.Prog()

	log.Println(tree)
	evalVisitor := EvalVisitor{}

	// result := evalVisitor.Visit(tree)
	evalVisitor.Visit(tree)

	log.Println(sourseFileName)

	// log.Println(result)
}
