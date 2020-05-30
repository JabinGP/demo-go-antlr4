// Code generated from Cal.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Cal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalParser.
type CalVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by CalParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by CalParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by CalParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by CalParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by CalParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by CalParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by CalParser#int.
	VisitInt(ctx *IntContext) interface{}
}
