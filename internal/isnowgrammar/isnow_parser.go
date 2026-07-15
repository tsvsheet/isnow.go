// Code generated from IsnowParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package isnowgrammar // IsnowParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type IsnowParser struct {
	*antlr.BaseParser
}

var IsnowParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func isnowparserParserInit() {
	staticData := &IsnowParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'>='", "'<='", "'>'", "'<'", "'/'", "':'", "'*'", "'!'", "','",
		"'-'", "'+'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "GE", "LE", "GT", "LT", "SLASH", "COLON", "STAR", "BANG", "COMMA",
		"DASH", "PLUS", "LBRACK", "RBRACK", "NUMBER", "NAME", "GSEP",
	}
	staticData.RuleNames = []string{
		"pattern", "bound", "boundOp", "exclusion", "spec", "group", "dateGroup",
		"timeGroup", "bareGroup", "field", "term", "incr", "atom", "qty",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 159, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 3, 0, 30, 8, 0, 1,
		0, 1, 0, 1, 0, 5, 0, 35, 8, 0, 10, 0, 12, 0, 38, 9, 0, 1, 0, 3, 0, 41,
		8, 0, 1, 0, 1, 0, 1, 1, 3, 1, 46, 8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 61, 8, 4, 10, 4, 12,
		4, 64, 9, 4, 1, 5, 1, 5, 1, 5, 3, 5, 69, 8, 5, 1, 6, 3, 6, 72, 8, 6, 1,
		6, 1, 6, 3, 6, 76, 8, 6, 4, 6, 78, 8, 6, 11, 6, 12, 6, 79, 1, 7, 3, 7,
		83, 8, 7, 1, 7, 1, 7, 3, 7, 87, 8, 7, 4, 7, 89, 8, 7, 11, 7, 12, 7, 90,
		1, 8, 1, 8, 1, 9, 3, 9, 96, 8, 9, 1, 9, 1, 9, 1, 9, 5, 9, 101, 8, 9, 10,
		9, 12, 9, 104, 9, 9, 1, 10, 3, 10, 107, 8, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 112, 8, 10, 1, 10, 3, 10, 115, 8, 10, 1, 10, 3, 10, 118, 8, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 125, 8, 11, 10, 11, 12, 11, 128,
		9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 137, 8,
		11, 10, 11, 12, 11, 140, 9, 11, 1, 11, 1, 11, 3, 11, 144, 8, 11, 1, 12,
		1, 12, 1, 12, 4, 12, 149, 8, 12, 11, 12, 12, 12, 150, 3, 12, 153, 8, 12,
		1, 13, 1, 13, 3, 13, 157, 8, 13, 1, 13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 0, 1, 1, 0, 1, 4, 171, 0, 29, 1, 0, 0, 0, 2,
		45, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0,
		10, 68, 1, 0, 0, 0, 12, 71, 1, 0, 0, 0, 14, 82, 1, 0, 0, 0, 16, 92, 1,
		0, 0, 0, 18, 95, 1, 0, 0, 0, 20, 117, 1, 0, 0, 0, 22, 143, 1, 0, 0, 0,
		24, 152, 1, 0, 0, 0, 26, 154, 1, 0, 0, 0, 28, 30, 5, 16, 0, 0, 29, 28,
		1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 36, 3, 8, 4, 0,
		32, 35, 3, 2, 1, 0, 33, 35, 3, 6, 3, 0, 34, 32, 1, 0, 0, 0, 34, 33, 1,
		0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37,
		40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 41, 5, 16, 0, 0, 40, 39, 1, 0,
		0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 5, 0, 0, 1, 43, 1,
		1, 0, 0, 0, 44, 46, 5, 16, 0, 0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 48, 3, 4, 2, 0, 48, 49, 3, 8, 4, 0, 49, 3, 1, 0,
		0, 0, 50, 51, 7, 0, 0, 0, 51, 5, 1, 0, 0, 0, 52, 53, 5, 16, 0, 0, 53, 54,
		5, 8, 0, 0, 54, 55, 5, 16, 0, 0, 55, 56, 3, 8, 4, 0, 56, 7, 1, 0, 0, 0,
		57, 62, 3, 10, 5, 0, 58, 59, 5, 16, 0, 0, 59, 61, 3, 10, 5, 0, 60, 58,
		1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0,
		63, 9, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 69, 3, 12, 6, 0, 66, 69, 3,
		14, 7, 0, 67, 69, 3, 16, 8, 0, 68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0,
		68, 67, 1, 0, 0, 0, 69, 11, 1, 0, 0, 0, 70, 72, 3, 18, 9, 0, 71, 70, 1,
		0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 77, 1, 0, 0, 0, 73, 75, 5, 5, 0, 0, 74,
		76, 3, 18, 9, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0,
		0, 0, 77, 73, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80,
		1, 0, 0, 0, 80, 13, 1, 0, 0, 0, 81, 83, 3, 18, 9, 0, 82, 81, 1, 0, 0, 0,
		82, 83, 1, 0, 0, 0, 83, 88, 1, 0, 0, 0, 84, 86, 5, 6, 0, 0, 85, 87, 3,
		18, 9, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88,
		84, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0,
		0, 91, 15, 1, 0, 0, 0, 92, 93, 3, 18, 9, 0, 93, 17, 1, 0, 0, 0, 94, 96,
		5, 8, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 102, 3, 20, 10, 0, 98, 99, 5, 9, 0, 0, 99, 101, 3, 20, 10, 0, 100,
		98, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1,
		0, 0, 0, 103, 19, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 107, 5, 10, 0,
		0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		111, 3, 24, 12, 0, 109, 110, 5, 10, 0, 0, 110, 112, 3, 24, 12, 0, 111,
		109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 115,
		3, 22, 11, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 118, 1,
		0, 0, 0, 116, 118, 3, 22, 11, 0, 117, 106, 1, 0, 0, 0, 117, 116, 1, 0,
		0, 0, 118, 21, 1, 0, 0, 0, 119, 120, 5, 11, 0, 0, 120, 121, 5, 12, 0, 0,
		121, 126, 3, 26, 13, 0, 122, 123, 5, 9, 0, 0, 123, 125, 3, 26, 13, 0, 124,
		122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127,
		1, 0, 0, 0, 127, 129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 130, 5, 13,
		0, 0, 130, 144, 1, 0, 0, 0, 131, 132, 5, 10, 0, 0, 132, 133, 5, 12, 0,
		0, 133, 138, 3, 26, 13, 0, 134, 135, 5, 9, 0, 0, 135, 137, 3, 26, 13, 0,
		136, 134, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142,
		5, 13, 0, 0, 142, 144, 1, 0, 0, 0, 143, 119, 1, 0, 0, 0, 143, 131, 1, 0,
		0, 0, 144, 23, 1, 0, 0, 0, 145, 153, 5, 7, 0, 0, 146, 153, 5, 15, 0, 0,
		147, 149, 3, 26, 13, 0, 148, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 145,
		1, 0, 0, 0, 152, 146, 1, 0, 0, 0, 152, 148, 1, 0, 0, 0, 153, 25, 1, 0,
		0, 0, 154, 156, 5, 14, 0, 0, 155, 157, 5, 15, 0, 0, 156, 155, 1, 0, 0,
		0, 156, 157, 1, 0, 0, 0, 157, 27, 1, 0, 0, 0, 25, 29, 34, 36, 40, 45, 62,
		68, 71, 75, 79, 82, 86, 90, 95, 102, 106, 111, 114, 117, 126, 138, 143,
		150, 152, 156,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// IsnowParserInit initializes any static state used to implement IsnowParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIsnowParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IsnowParserInit() {
	staticData := &IsnowParserParserStaticData
	staticData.once.Do(isnowparserParserInit)
}

// NewIsnowParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIsnowParser(input antlr.TokenStream) *IsnowParser {
	IsnowParserInit()
	this := new(IsnowParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &IsnowParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "IsnowParser.g4"

	return this
}

// IsnowParser tokens.
const (
	IsnowParserEOF    = antlr.TokenEOF
	IsnowParserGE     = 1
	IsnowParserLE     = 2
	IsnowParserGT     = 3
	IsnowParserLT     = 4
	IsnowParserSLASH  = 5
	IsnowParserCOLON  = 6
	IsnowParserSTAR   = 7
	IsnowParserBANG   = 8
	IsnowParserCOMMA  = 9
	IsnowParserDASH   = 10
	IsnowParserPLUS   = 11
	IsnowParserLBRACK = 12
	IsnowParserRBRACK = 13
	IsnowParserNUMBER = 14
	IsnowParserNAME   = 15
	IsnowParserGSEP   = 16
)

// IsnowParser rules.
const (
	IsnowParserRULE_pattern   = 0
	IsnowParserRULE_bound     = 1
	IsnowParserRULE_boundOp   = 2
	IsnowParserRULE_exclusion = 3
	IsnowParserRULE_spec      = 4
	IsnowParserRULE_group     = 5
	IsnowParserRULE_dateGroup = 6
	IsnowParserRULE_timeGroup = 7
	IsnowParserRULE_bareGroup = 8
	IsnowParserRULE_field     = 9
	IsnowParserRULE_term      = 10
	IsnowParserRULE_incr      = 11
	IsnowParserRULE_atom      = 12
	IsnowParserRULE_qty       = 13
)

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Spec() ISpecContext
	EOF() antlr.TerminalNode
	AllGSEP() []antlr.TerminalNode
	GSEP(i int) antlr.TerminalNode
	AllBound() []IBoundContext
	Bound(i int) IBoundContext
	AllExclusion() []IExclusionContext
	Exclusion(i int) IExclusionContext

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_pattern
	return p
}

func InitEmptyPatternContext(p *PatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_pattern
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) Spec() ISpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *PatternContext) EOF() antlr.TerminalNode {
	return s.GetToken(IsnowParserEOF, 0)
}

func (s *PatternContext) AllGSEP() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserGSEP)
}

func (s *PatternContext) GSEP(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserGSEP, i)
}

func (s *PatternContext) AllBound() []IBoundContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoundContext); ok {
			len++
		}
	}

	tst := make([]IBoundContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoundContext); ok {
			tst[i] = t.(IBoundContext)
			i++
		}
	}

	return tst
}

func (s *PatternContext) Bound(i int) IBoundContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoundContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoundContext)
}

func (s *PatternContext) AllExclusion() []IExclusionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExclusionContext); ok {
			len++
		}
	}

	tst := make([]IExclusionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExclusionContext); ok {
			tst[i] = t.(IExclusionContext)
			i++
		}
	}

	return tst
}

func (s *PatternContext) Exclusion(i int) IExclusionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExclusionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExclusionContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (s *PatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IsnowParserRULE_pattern)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == IsnowParserGSEP {
		{
			p.SetState(28)
			p.Match(IsnowParserGSEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(31)
		p.Spec()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(32)
					p.Bound()
				}

			case 2:
				{
					p.SetState(33)
					p.Exclusion()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == IsnowParserGSEP {
		{
			p.SetState(39)
			p.Match(IsnowParserGSEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(42)
		p.Match(IsnowParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoundContext is an interface to support dynamic dispatch.
type IBoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BoundOp() IBoundOpContext
	Spec() ISpecContext
	GSEP() antlr.TerminalNode

	// IsBoundContext differentiates from other interfaces.
	IsBoundContext()
}

type BoundContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundContext() *BoundContext {
	var p = new(BoundContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_bound
	return p
}

func InitEmptyBoundContext(p *BoundContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_bound
}

func (*BoundContext) IsBoundContext() {}

func NewBoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundContext {
	var p = new(BoundContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_bound

	return p
}

func (s *BoundContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundContext) BoundOp() IBoundOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoundOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoundOpContext)
}

func (s *BoundContext) Spec() ISpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *BoundContext) GSEP() antlr.TerminalNode {
	return s.GetToken(IsnowParserGSEP, 0)
}

func (s *BoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterBound(s)
	}
}

func (s *BoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitBound(s)
	}
}

func (s *BoundContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitBound(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Bound() (localctx IBoundContext) {
	localctx = NewBoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IsnowParserRULE_bound)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == IsnowParserGSEP {
		{
			p.SetState(44)
			p.Match(IsnowParserGSEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(47)
		p.BoundOp()
	}
	{
		p.SetState(48)
		p.Spec()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoundOpContext is an interface to support dynamic dispatch.
type IBoundOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GE() antlr.TerminalNode
	GT() antlr.TerminalNode
	LE() antlr.TerminalNode
	LT() antlr.TerminalNode

	// IsBoundOpContext differentiates from other interfaces.
	IsBoundOpContext()
}

type BoundOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundOpContext() *BoundOpContext {
	var p = new(BoundOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_boundOp
	return p
}

func InitEmptyBoundOpContext(p *BoundOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_boundOp
}

func (*BoundOpContext) IsBoundOpContext() {}

func NewBoundOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundOpContext {
	var p = new(BoundOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_boundOp

	return p
}

func (s *BoundOpContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundOpContext) GE() antlr.TerminalNode {
	return s.GetToken(IsnowParserGE, 0)
}

func (s *BoundOpContext) GT() antlr.TerminalNode {
	return s.GetToken(IsnowParserGT, 0)
}

func (s *BoundOpContext) LE() antlr.TerminalNode {
	return s.GetToken(IsnowParserLE, 0)
}

func (s *BoundOpContext) LT() antlr.TerminalNode {
	return s.GetToken(IsnowParserLT, 0)
}

func (s *BoundOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterBoundOp(s)
	}
}

func (s *BoundOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitBoundOp(s)
	}
}

func (s *BoundOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitBoundOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) BoundOp() (localctx IBoundOpContext) {
	localctx = NewBoundOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, IsnowParserRULE_boundOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExclusionContext is an interface to support dynamic dispatch.
type IExclusionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGSEP() []antlr.TerminalNode
	GSEP(i int) antlr.TerminalNode
	BANG() antlr.TerminalNode
	Spec() ISpecContext

	// IsExclusionContext differentiates from other interfaces.
	IsExclusionContext()
}

type ExclusionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExclusionContext() *ExclusionContext {
	var p = new(ExclusionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_exclusion
	return p
}

func InitEmptyExclusionContext(p *ExclusionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_exclusion
}

func (*ExclusionContext) IsExclusionContext() {}

func NewExclusionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusionContext {
	var p = new(ExclusionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_exclusion

	return p
}

func (s *ExclusionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExclusionContext) AllGSEP() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserGSEP)
}

func (s *ExclusionContext) GSEP(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserGSEP, i)
}

func (s *ExclusionContext) BANG() antlr.TerminalNode {
	return s.GetToken(IsnowParserBANG, 0)
}

func (s *ExclusionContext) Spec() ISpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *ExclusionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExclusionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterExclusion(s)
	}
}

func (s *ExclusionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitExclusion(s)
	}
}

func (s *ExclusionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitExclusion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Exclusion() (localctx IExclusionContext) {
	localctx = NewExclusionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, IsnowParserRULE_exclusion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(IsnowParserGSEP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(IsnowParserBANG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(IsnowParserGSEP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Spec()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGroup() []IGroupContext
	Group(i int) IGroupContext
	AllGSEP() []antlr.TerminalNode
	GSEP(i int) antlr.TerminalNode

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_spec
	return p
}

func InitEmptySpecContext(p *SpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_spec
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) AllGroup() []IGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupContext); ok {
			len++
		}
	}

	tst := make([]IGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupContext); ok {
			tst[i] = t.(IGroupContext)
			i++
		}
	}

	return tst
}

func (s *SpecContext) Group(i int) IGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *SpecContext) AllGSEP() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserGSEP)
}

func (s *SpecContext) GSEP(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserGSEP, i)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterSpec(s)
	}
}

func (s *SpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitSpec(s)
	}
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Spec() (localctx ISpecContext) {
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IsnowParserRULE_spec)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Group()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(58)
				p.Match(IsnowParserGSEP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(59)
				p.Group()
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DateGroup() IDateGroupContext
	TimeGroup() ITimeGroupContext
	BareGroup() IBareGroupContext

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_group
	return p
}

func InitEmptyGroupContext(p *GroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_group
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) DateGroup() IDateGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateGroupContext)
}

func (s *GroupContext) TimeGroup() ITimeGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeGroupContext)
}

func (s *GroupContext) BareGroup() IBareGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBareGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBareGroupContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IsnowParserRULE_group)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.DateGroup()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.TimeGroup()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.BareGroup()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateGroupContext is an interface to support dynamic dispatch.
type IDateGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode

	// IsDateGroupContext differentiates from other interfaces.
	IsDateGroupContext()
}

type DateGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateGroupContext() *DateGroupContext {
	var p = new(DateGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_dateGroup
	return p
}

func InitEmptyDateGroupContext(p *DateGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_dateGroup
}

func (*DateGroupContext) IsDateGroupContext() {}

func NewDateGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateGroupContext {
	var p = new(DateGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_dateGroup

	return p
}

func (s *DateGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *DateGroupContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *DateGroupContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *DateGroupContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserSLASH)
}

func (s *DateGroupContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserSLASH, i)
}

func (s *DateGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterDateGroup(s)
	}
}

func (s *DateGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitDateGroup(s)
	}
}

func (s *DateGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitDateGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) DateGroup() (localctx IDateGroupContext) {
	localctx = NewDateGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, IsnowParserRULE_dateGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52608) != 0 {
		{
			p.SetState(70)
			p.Field()
		}

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == IsnowParserSLASH {
		{
			p.SetState(73)
			p.Match(IsnowParserSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52608) != 0 {
			{
				p.SetState(74)
				p.Field()
			}

		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeGroupContext is an interface to support dynamic dispatch.
type ITimeGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode

	// IsTimeGroupContext differentiates from other interfaces.
	IsTimeGroupContext()
}

type TimeGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeGroupContext() *TimeGroupContext {
	var p = new(TimeGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_timeGroup
	return p
}

func InitEmptyTimeGroupContext(p *TimeGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_timeGroup
}

func (*TimeGroupContext) IsTimeGroupContext() {}

func NewTimeGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeGroupContext {
	var p = new(TimeGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_timeGroup

	return p
}

func (s *TimeGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeGroupContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TimeGroupContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TimeGroupContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserCOLON)
}

func (s *TimeGroupContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserCOLON, i)
}

func (s *TimeGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterTimeGroup(s)
	}
}

func (s *TimeGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitTimeGroup(s)
	}
}

func (s *TimeGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitTimeGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) TimeGroup() (localctx ITimeGroupContext) {
	localctx = NewTimeGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, IsnowParserRULE_timeGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52608) != 0 {
		{
			p.SetState(81)
			p.Field()
		}

	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == IsnowParserCOLON {
		{
			p.SetState(84)
			p.Match(IsnowParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52608) != 0 {
			{
				p.SetState(85)
				p.Field()
			}

		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBareGroupContext is an interface to support dynamic dispatch.
type IBareGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext

	// IsBareGroupContext differentiates from other interfaces.
	IsBareGroupContext()
}

type BareGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBareGroupContext() *BareGroupContext {
	var p = new(BareGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_bareGroup
	return p
}

func InitEmptyBareGroupContext(p *BareGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_bareGroup
}

func (*BareGroupContext) IsBareGroupContext() {}

func NewBareGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BareGroupContext {
	var p = new(BareGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_bareGroup

	return p
}

func (s *BareGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *BareGroupContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *BareGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BareGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BareGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterBareGroup(s)
	}
}

func (s *BareGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitBareGroup(s)
	}
}

func (s *BareGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitBareGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) BareGroup() (localctx IBareGroupContext) {
	localctx = NewBareGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IsnowParserRULE_bareGroup)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Field()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	BANG() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FieldContext) BANG() antlr.TerminalNode {
	return s.GetToken(IsnowParserBANG, 0)
}

func (s *FieldContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserCOMMA)
}

func (s *FieldContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserCOMMA, i)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, IsnowParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == IsnowParserBANG {
		{
			p.SetState(94)
			p.Match(IsnowParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(97)
		p.Term()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == IsnowParserCOMMA {
		{
			p.SetState(98)
			p.Match(IsnowParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Term()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAtom() []IAtomContext
	Atom(i int) IAtomContext
	AllDASH() []antlr.TerminalNode
	DASH(i int) antlr.TerminalNode
	Incr() IIncrContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllAtom() []IAtomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtomContext); ok {
			len++
		}
	}

	tst := make([]IAtomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtomContext); ok {
			tst[i] = t.(IAtomContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Atom(i int) IAtomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *TermContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserDASH)
}

func (s *TermContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserDASH, i)
}

func (s *TermContext) Incr() IIncrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncrContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, IsnowParserRULE_term)
	var _la int

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == IsnowParserDASH {
			{
				p.SetState(105)
				p.Match(IsnowParserDASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(108)
			p.Atom()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(109)
				p.Match(IsnowParserDASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(110)
				p.Atom()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == IsnowParserDASH || _la == IsnowParserPLUS {
			{
				p.SetState(113)
				p.Incr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Incr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncrContext is an interface to support dynamic dispatch.
type IIncrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	AllQty() []IQtyContext
	Qty(i int) IQtyContext
	RBRACK() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	DASH() antlr.TerminalNode

	// IsIncrContext differentiates from other interfaces.
	IsIncrContext()
}

type IncrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncrContext() *IncrContext {
	var p = new(IncrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_incr
	return p
}

func InitEmptyIncrContext(p *IncrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_incr
}

func (*IncrContext) IsIncrContext() {}

func NewIncrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncrContext {
	var p = new(IncrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_incr

	return p
}

func (s *IncrContext) GetParser() antlr.Parser { return s.parser }

func (s *IncrContext) PLUS() antlr.TerminalNode {
	return s.GetToken(IsnowParserPLUS, 0)
}

func (s *IncrContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(IsnowParserLBRACK, 0)
}

func (s *IncrContext) AllQty() []IQtyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQtyContext); ok {
			len++
		}
	}

	tst := make([]IQtyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQtyContext); ok {
			tst[i] = t.(IQtyContext)
			i++
		}
	}

	return tst
}

func (s *IncrContext) Qty(i int) IQtyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQtyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQtyContext)
}

func (s *IncrContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(IsnowParserRBRACK, 0)
}

func (s *IncrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IsnowParserCOMMA)
}

func (s *IncrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IsnowParserCOMMA, i)
}

func (s *IncrContext) DASH() antlr.TerminalNode {
	return s.GetToken(IsnowParserDASH, 0)
}

func (s *IncrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterIncr(s)
	}
}

func (s *IncrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitIncr(s)
	}
}

func (s *IncrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitIncr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Incr() (localctx IIncrContext) {
	localctx = NewIncrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, IsnowParserRULE_incr)
	var _la int

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case IsnowParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(IsnowParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(IsnowParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Qty()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == IsnowParserCOMMA {
			{
				p.SetState(122)
				p.Match(IsnowParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(123)
				p.Qty()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(IsnowParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IsnowParserDASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Match(IsnowParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(IsnowParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Qty()
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == IsnowParserCOMMA {
			{
				p.SetState(134)
				p.Match(IsnowParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(135)
				p.Qty()
			}

			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(141)
			p.Match(IsnowParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	NAME() antlr.TerminalNode
	AllQty() []IQtyContext
	Qty(i int) IQtyContext

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) STAR() antlr.TerminalNode {
	return s.GetToken(IsnowParserSTAR, 0)
}

func (s *AtomContext) NAME() antlr.TerminalNode {
	return s.GetToken(IsnowParserNAME, 0)
}

func (s *AtomContext) AllQty() []IQtyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQtyContext); ok {
			len++
		}
	}

	tst := make([]IQtyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQtyContext); ok {
			tst[i] = t.(IQtyContext)
			i++
		}
	}

	return tst
}

func (s *AtomContext) Qty(i int) IQtyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQtyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQtyContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, IsnowParserRULE_atom)
	var _la int

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case IsnowParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(IsnowParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IsnowParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(IsnowParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IsnowParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == IsnowParserNUMBER {
			{
				p.SetState(147)
				p.Qty()
			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQtyContext is an interface to support dynamic dispatch.
type IQtyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsQtyContext differentiates from other interfaces.
	IsQtyContext()
}

type QtyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQtyContext() *QtyContext {
	var p = new(QtyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_qty
	return p
}

func InitEmptyQtyContext(p *QtyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IsnowParserRULE_qty
}

func (*QtyContext) IsQtyContext() {}

func NewQtyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QtyContext {
	var p = new(QtyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IsnowParserRULE_qty

	return p
}

func (s *QtyContext) GetParser() antlr.Parser { return s.parser }

func (s *QtyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(IsnowParserNUMBER, 0)
}

func (s *QtyContext) NAME() antlr.TerminalNode {
	return s.GetToken(IsnowParserNAME, 0)
}

func (s *QtyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QtyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QtyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.EnterQty(s)
	}
}

func (s *QtyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IsnowParserListener); ok {
		listenerT.ExitQty(s)
	}
}

func (s *QtyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IsnowParserVisitor:
		return t.VisitQty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IsnowParser) Qty() (localctx IQtyContext) {
	localctx = NewQtyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, IsnowParserRULE_qty)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(IsnowParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == IsnowParserNAME {
		{
			p.SetState(155)
			p.Match(IsnowParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
