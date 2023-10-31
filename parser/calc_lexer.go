// Code generated from f://mario//repositories//mycalculator//parser//Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

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

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalcLexerLexerStaticData struct {
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

func calclexerLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'\\n'", "'='", "'^'", "'*'", "'/'", "'+'", "'-'", "", "", "'('",
		"')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "MUL", "DIV", "ADD", "SUB", "NUMBER", "VARIABLE_NAME",
		"OPEN_BRACKET", "CLOSE_BRACKET", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "MUL", "DIV", "ADD", "SUB", "NUMBER", "VARIABLE_NAME",
		"OPEN_BRACKET", "CLOSE_BRACKET", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 62, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 4, 7, 41, 8, 7, 11, 7, 12, 7,
		42, 1, 8, 1, 8, 5, 8, 47, 8, 8, 10, 8, 12, 8, 50, 9, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 4, 11, 57, 8, 11, 11, 11, 12, 11, 58, 1, 11, 1, 11, 0,
		0, 12, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 1, 0, 4, 1, 0, 48, 57, 1, 0, 97, 122, 2, 0, 48, 57, 97,
		122, 3, 0, 9, 10, 13, 13, 32, 32, 64, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 27, 1,
		0, 0, 0, 5, 29, 1, 0, 0, 0, 7, 31, 1, 0, 0, 0, 9, 33, 1, 0, 0, 0, 11, 35,
		1, 0, 0, 0, 13, 37, 1, 0, 0, 0, 15, 40, 1, 0, 0, 0, 17, 44, 1, 0, 0, 0,
		19, 51, 1, 0, 0, 0, 21, 53, 1, 0, 0, 0, 23, 56, 1, 0, 0, 0, 25, 26, 5,
		10, 0, 0, 26, 2, 1, 0, 0, 0, 27, 28, 5, 61, 0, 0, 28, 4, 1, 0, 0, 0, 29,
		30, 5, 94, 0, 0, 30, 6, 1, 0, 0, 0, 31, 32, 5, 42, 0, 0, 32, 8, 1, 0, 0,
		0, 33, 34, 5, 47, 0, 0, 34, 10, 1, 0, 0, 0, 35, 36, 5, 43, 0, 0, 36, 12,
		1, 0, 0, 0, 37, 38, 5, 45, 0, 0, 38, 14, 1, 0, 0, 0, 39, 41, 7, 0, 0, 0,
		40, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1,
		0, 0, 0, 43, 16, 1, 0, 0, 0, 44, 48, 7, 1, 0, 0, 45, 47, 7, 2, 0, 0, 46,
		45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0,
		0, 49, 18, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 52, 5, 40, 0, 0, 52, 20,
		1, 0, 0, 0, 53, 54, 5, 41, 0, 0, 54, 22, 1, 0, 0, 0, 55, 57, 7, 3, 0, 0,
		56, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 6, 11, 0, 0, 61, 24, 1, 0, 0, 0, 4,
		0, 42, 48, 58, 1, 6, 0, 0,
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

// CalcLexerInit initializes any static state used to implement CalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalcLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerT__0          = 1
	CalcLexerT__1          = 2
	CalcLexerT__2          = 3
	CalcLexerMUL           = 4
	CalcLexerDIV           = 5
	CalcLexerADD           = 6
	CalcLexerSUB           = 7
	CalcLexerNUMBER        = 8
	CalcLexerVARIABLE_NAME = 9
	CalcLexerOPEN_BRACKET  = 10
	CalcLexerCLOSE_BRACKET = 11
	CalcLexerWHITESPACE    = 12
)
