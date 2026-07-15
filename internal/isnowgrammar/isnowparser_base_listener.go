// Code generated from IsnowParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package isnowgrammar // IsnowParser
import "github.com/antlr4-go/antlr/v4"

// BaseIsnowParserListener is a complete listener for a parse tree produced by IsnowParser.
type BaseIsnowParserListener struct{}

var _ IsnowParserListener = &BaseIsnowParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIsnowParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIsnowParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIsnowParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIsnowParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseIsnowParserListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseIsnowParserListener) ExitPattern(ctx *PatternContext) {}

// EnterBound is called when production bound is entered.
func (s *BaseIsnowParserListener) EnterBound(ctx *BoundContext) {}

// ExitBound is called when production bound is exited.
func (s *BaseIsnowParserListener) ExitBound(ctx *BoundContext) {}

// EnterBoundOp is called when production boundOp is entered.
func (s *BaseIsnowParserListener) EnterBoundOp(ctx *BoundOpContext) {}

// ExitBoundOp is called when production boundOp is exited.
func (s *BaseIsnowParserListener) ExitBoundOp(ctx *BoundOpContext) {}

// EnterExclusion is called when production exclusion is entered.
func (s *BaseIsnowParserListener) EnterExclusion(ctx *ExclusionContext) {}

// ExitExclusion is called when production exclusion is exited.
func (s *BaseIsnowParserListener) ExitExclusion(ctx *ExclusionContext) {}

// EnterSpec is called when production spec is entered.
func (s *BaseIsnowParserListener) EnterSpec(ctx *SpecContext) {}

// ExitSpec is called when production spec is exited.
func (s *BaseIsnowParserListener) ExitSpec(ctx *SpecContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseIsnowParserListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseIsnowParserListener) ExitGroup(ctx *GroupContext) {}

// EnterDateGroup is called when production dateGroup is entered.
func (s *BaseIsnowParserListener) EnterDateGroup(ctx *DateGroupContext) {}

// ExitDateGroup is called when production dateGroup is exited.
func (s *BaseIsnowParserListener) ExitDateGroup(ctx *DateGroupContext) {}

// EnterTimeGroup is called when production timeGroup is entered.
func (s *BaseIsnowParserListener) EnterTimeGroup(ctx *TimeGroupContext) {}

// ExitTimeGroup is called when production timeGroup is exited.
func (s *BaseIsnowParserListener) ExitTimeGroup(ctx *TimeGroupContext) {}

// EnterBareGroup is called when production bareGroup is entered.
func (s *BaseIsnowParserListener) EnterBareGroup(ctx *BareGroupContext) {}

// ExitBareGroup is called when production bareGroup is exited.
func (s *BaseIsnowParserListener) ExitBareGroup(ctx *BareGroupContext) {}

// EnterField is called when production field is entered.
func (s *BaseIsnowParserListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseIsnowParserListener) ExitField(ctx *FieldContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseIsnowParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseIsnowParserListener) ExitTerm(ctx *TermContext) {}

// EnterIncr is called when production incr is entered.
func (s *BaseIsnowParserListener) EnterIncr(ctx *IncrContext) {}

// ExitIncr is called when production incr is exited.
func (s *BaseIsnowParserListener) ExitIncr(ctx *IncrContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseIsnowParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseIsnowParserListener) ExitAtom(ctx *AtomContext) {}

// EnterQty is called when production qty is entered.
func (s *BaseIsnowParserListener) EnterQty(ctx *QtyContext) {}

// ExitQty is called when production qty is exited.
func (s *BaseIsnowParserListener) ExitQty(ctx *QtyContext) {}
