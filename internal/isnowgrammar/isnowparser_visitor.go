// Code generated from IsnowParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package isnowgrammar // IsnowParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by IsnowParser.
type IsnowParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by IsnowParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by IsnowParser#bound.
	VisitBound(ctx *BoundContext) interface{}

	// Visit a parse tree produced by IsnowParser#boundOp.
	VisitBoundOp(ctx *BoundOpContext) interface{}

	// Visit a parse tree produced by IsnowParser#exclusion.
	VisitExclusion(ctx *ExclusionContext) interface{}

	// Visit a parse tree produced by IsnowParser#spec.
	VisitSpec(ctx *SpecContext) interface{}

	// Visit a parse tree produced by IsnowParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by IsnowParser#dateGroup.
	VisitDateGroup(ctx *DateGroupContext) interface{}

	// Visit a parse tree produced by IsnowParser#timeGroup.
	VisitTimeGroup(ctx *TimeGroupContext) interface{}

	// Visit a parse tree produced by IsnowParser#bareGroup.
	VisitBareGroup(ctx *BareGroupContext) interface{}

	// Visit a parse tree produced by IsnowParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by IsnowParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by IsnowParser#incr.
	VisitIncr(ctx *IncrContext) interface{}

	// Visit a parse tree produced by IsnowParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by IsnowParser#qty.
	VisitQty(ctx *QtyContext) interface{}
}
