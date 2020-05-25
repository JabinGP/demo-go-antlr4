package main

import (
	"log"
	"os"

	"github.com/JabinGP/demo-go-antlr4/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseJSONListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	log.Println(ctx.GetText())
}

func main() {

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewJSONLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewJSONParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Json()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
