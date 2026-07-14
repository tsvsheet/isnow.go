// Code generated from IsnowParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package isnowgrammar // IsnowParser
import "github.com/antlr4-go/antlr/v4"

type BaseIsnowParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseIsnowParserVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitBound(ctx *BoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitBoundOp(ctx *BoundOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitSpec(ctx *SpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitDateGroup(ctx *DateGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitTimeGroup(ctx *TimeGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitBareGroup(ctx *BareGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitIncr(ctx *IncrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIsnowParserVisitor) VisitQty(ctx *QtyContext) interface{} {
	return v.VisitChildren(ctx)
}
