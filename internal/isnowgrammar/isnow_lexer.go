// Code generated from IsnowLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package isnowgrammar

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type IsnowLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var IsnowLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func isnowlexerLexerInit() {
	staticData := &IsnowLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'>='", "'<='", "'>'", "'<'", "'/'", "':'", "'*'", "'!'", "','",
		"'-'", "'+'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "GE", "LE", "GT", "LT", "SLASH", "COLON", "STAR", "BANG", "COMMA",
		"DASH", "PLUS", "LBRACK", "RBRACK", "NUMBER", "NAME", "GSEP",
	}
	staticData.RuleNames = []string{
		"GE", "LE", "GT", "LT", "SLASH", "COLON", "STAR", "BANG", "COMMA", "DASH",
		"PLUS", "LBRACK", "RBRACK", "NUMBER", "NAME", "GSEP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 76, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 4, 13, 63, 8, 13, 11, 13, 12,
		13, 64, 1, 14, 4, 14, 68, 8, 14, 11, 14, 12, 14, 69, 1, 15, 4, 15, 73,
		8, 15, 11, 15, 12, 15, 74, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 1, 0, 3, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 4, 0, 9, 9, 32, 32,
		46, 46, 95, 95, 78, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 36, 1, 0, 0, 0, 5,
		39, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 43, 1, 0, 0, 0, 11, 45, 1, 0, 0,
		0, 13, 47, 1, 0, 0, 0, 15, 49, 1, 0, 0, 0, 17, 51, 1, 0, 0, 0, 19, 53,
		1, 0, 0, 0, 21, 55, 1, 0, 0, 0, 23, 57, 1, 0, 0, 0, 25, 59, 1, 0, 0, 0,
		27, 62, 1, 0, 0, 0, 29, 67, 1, 0, 0, 0, 31, 72, 1, 0, 0, 0, 33, 34, 5,
		62, 0, 0, 34, 35, 5, 61, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 60, 0, 0,
		37, 38, 5, 61, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 62, 0, 0, 40, 6, 1,
		0, 0, 0, 41, 42, 5, 60, 0, 0, 42, 8, 1, 0, 0, 0, 43, 44, 5, 47, 0, 0, 44,
		10, 1, 0, 0, 0, 45, 46, 5, 58, 0, 0, 46, 12, 1, 0, 0, 0, 47, 48, 5, 42,
		0, 0, 48, 14, 1, 0, 0, 0, 49, 50, 5, 33, 0, 0, 50, 16, 1, 0, 0, 0, 51,
		52, 5, 44, 0, 0, 52, 18, 1, 0, 0, 0, 53, 54, 5, 45, 0, 0, 54, 20, 1, 0,
		0, 0, 55, 56, 5, 43, 0, 0, 56, 22, 1, 0, 0, 0, 57, 58, 5, 91, 0, 0, 58,
		24, 1, 0, 0, 0, 59, 60, 5, 93, 0, 0, 60, 26, 1, 0, 0, 0, 61, 63, 7, 0,
		0, 0, 62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65,
		1, 0, 0, 0, 65, 28, 1, 0, 0, 0, 66, 68, 7, 1, 0, 0, 67, 66, 1, 0, 0, 0,
		68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 30, 1,
		0, 0, 0, 71, 73, 7, 2, 0, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 32, 1, 0, 0, 0, 4, 0, 64, 69, 74,
		0,
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

// IsnowLexerInit initializes any static state used to implement IsnowLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewIsnowLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func IsnowLexerInit() {
	staticData := &IsnowLexerLexerStaticData
	staticData.once.Do(isnowlexerLexerInit)
}

// NewIsnowLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewIsnowLexer(input antlr.CharStream) *IsnowLexer {
	IsnowLexerInit()
	l := new(IsnowLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &IsnowLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "IsnowLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// IsnowLexer tokens.
const (
	IsnowLexerGE     = 1
	IsnowLexerLE     = 2
	IsnowLexerGT     = 3
	IsnowLexerLT     = 4
	IsnowLexerSLASH  = 5
	IsnowLexerCOLON  = 6
	IsnowLexerSTAR   = 7
	IsnowLexerBANG   = 8
	IsnowLexerCOMMA  = 9
	IsnowLexerDASH   = 10
	IsnowLexerPLUS   = 11
	IsnowLexerLBRACK = 12
	IsnowLexerRBRACK = 13
	IsnowLexerNUMBER = 14
	IsnowLexerNAME   = 15
	IsnowLexerGSEP   = 16
)
