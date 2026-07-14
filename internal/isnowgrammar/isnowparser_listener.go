// Code generated from IsnowParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package isnowgrammar // IsnowParser
import "github.com/antlr4-go/antlr/v4"

// IsnowParserListener is a complete listener for a parse tree produced by IsnowParser.
type IsnowParserListener interface {
	antlr.ParseTreeListener

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterBound is called when entering the bound production.
	EnterBound(c *BoundContext)

	// EnterBoundOp is called when entering the boundOp production.
	EnterBoundOp(c *BoundOpContext)

	// EnterSpec is called when entering the spec production.
	EnterSpec(c *SpecContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterDateGroup is called when entering the dateGroup production.
	EnterDateGroup(c *DateGroupContext)

	// EnterTimeGroup is called when entering the timeGroup production.
	EnterTimeGroup(c *TimeGroupContext)

	// EnterBareGroup is called when entering the bareGroup production.
	EnterBareGroup(c *BareGroupContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterIncr is called when entering the incr production.
	EnterIncr(c *IncrContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterQty is called when entering the qty production.
	EnterQty(c *QtyContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitBound is called when exiting the bound production.
	ExitBound(c *BoundContext)

	// ExitBoundOp is called when exiting the boundOp production.
	ExitBoundOp(c *BoundOpContext)

	// ExitSpec is called when exiting the spec production.
	ExitSpec(c *SpecContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitDateGroup is called when exiting the dateGroup production.
	ExitDateGroup(c *DateGroupContext)

	// ExitTimeGroup is called when exiting the timeGroup production.
	ExitTimeGroup(c *TimeGroupContext)

	// ExitBareGroup is called when exiting the bareGroup production.
	ExitBareGroup(c *BareGroupContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitIncr is called when exiting the incr production.
	ExitIncr(c *IncrContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitQty is called when exiting the qty production.
	ExitQty(c *QtyContext)
}
