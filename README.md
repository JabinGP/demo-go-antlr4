# demo-go-antlr4

由于 Go 语言无法支持方法名调用方法，官方的 runtime 的 Visitor 方法也尚未实现，代码目前无法运行。

以下代码来自官方

```go
type BaseParseTreeVisitor struct{}

func (v *BaseParseTreeVisitor) Visit(tree ParseTree) interface{}            { return nil }
func (v *BaseParseTreeVisitor) VisitChildren(node RuleNode) interface{}     { return nil }
func (v *BaseParseTreeVisitor) VisitTerminal(node TerminalNode) interface{} { return nil }
func (v *BaseParseTreeVisitor) VisitErrorNode(node ErrorNode) interface{}   { return nil }

// TODO
//func (this ParseTreeVisitor) Visit(ctx) {
//	if (Utils.isArray(ctx)) {
//		self := this
//		return ctx.map(function(child) { return VisitAtom(self, child)})
//	} else {
//		return VisitAtom(this, ctx)
//	}
//}
//
//func VisitAtom(Visitor, ctx) {
//	if (ctx.parser == nil) { //is terminal
//		return
//	}
//
//	name := ctx.parser.ruleNames[ctx.ruleIndex]
//	funcName := "Visit" + Utils.titleCase(name)
//
//	return Visitor[funcName](ctx)
//}
```
