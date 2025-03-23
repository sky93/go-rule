// Code generated from internal/antlr4/JsonQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JsonQuery

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

type JsonQueryParser struct {
	*antlr.BaseParser
}

var JsonQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonqueryParserInit() {
	staticData := &JsonQueryParserStaticData
	staticData.LiteralNames = []string{
		"", "'pr'", "'[f64]'", "'[i64]'", "'[ui64]'", "'[i]'", "'[ui]'", "'[i32]'",
		"'[ui32]'", "'[d]'", "'[s]'", "'[f32]'", "'.'", "'-'", "'['", "']'",
		"", "", "", "'null'", "", "", "", "", "", "", "", "", "", "", "'dgt'",
		"'deq'", "", "", "", "", "", "'('", "')'", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NOT",
		"LOGICAL_OPERATOR", "BOOLEAN", "NULL", "IN", "EQ", "NE", "GT", "LT",
		"GE", "LE", "CO", "SW", "EW", "DGT", "DEQ", "ATTRNAME", "VERSION", "STRING",
		"DOUBLE", "INT", "LPAREN", "RPAREN", "EXP", "NEWLINE", "COMMA", "SP",
	}
	staticData.RuleNames = []string{
		"root", "query", "attrPath", "typeAnnotation", "functionCall", "argList",
		"subAttr", "typedValue", "value", "listStrings", "subListOfStrings",
		"listDoubles", "subListOfDoubles", "listInts", "subListOfInts",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 162, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 3, 1, 39, 8, 1, 1, 1, 1, 1, 3,
		1, 43, 8, 1, 1, 1, 1, 1, 3, 1, 47, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 57, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 64,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 71, 8, 1, 10, 1, 12, 1, 74, 9,
		1, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 2, 3, 2, 81, 8, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 3, 4, 88, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 95, 8,
		5, 10, 5, 12, 5, 98, 9, 5, 1, 6, 1, 6, 1, 6, 1, 7, 3, 7, 104, 8, 7, 1,
		7, 1, 7, 3, 7, 108, 8, 7, 1, 7, 1, 7, 3, 7, 112, 8, 7, 1, 7, 3, 7, 115,
		8, 7, 1, 7, 1, 7, 3, 7, 119, 8, 7, 3, 7, 121, 8, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 130, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 140, 8, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 150, 8, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 160, 8, 14, 1, 14, 0, 1, 2, 15, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 2, 1, 0, 20, 30, 1, 0,
		2, 11, 174, 0, 30, 1, 0, 0, 0, 2, 63, 1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6,
		82, 1, 0, 0, 0, 8, 84, 1, 0, 0, 0, 10, 91, 1, 0, 0, 0, 12, 99, 1, 0, 0,
		0, 14, 120, 1, 0, 0, 0, 16, 129, 1, 0, 0, 0, 18, 131, 1, 0, 0, 0, 20, 139,
		1, 0, 0, 0, 22, 141, 1, 0, 0, 0, 24, 149, 1, 0, 0, 0, 26, 151, 1, 0, 0,
		0, 28, 159, 1, 0, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 0, 0, 1, 32, 1,
		1, 0, 0, 0, 33, 35, 6, 1, -1, 0, 34, 36, 5, 16, 0, 0, 35, 34, 1, 0, 0,
		0, 35, 36, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 39, 5, 42, 0, 0, 38, 37,
		1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 5, 37, 0, 0,
		41, 43, 5, 42, 0, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1,
		0, 0, 0, 44, 46, 3, 2, 1, 0, 45, 47, 5, 42, 0, 0, 46, 45, 1, 0, 0, 0, 46,
		47, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 5, 38, 0, 0, 49, 64, 1, 0,
		0, 0, 50, 51, 3, 4, 2, 0, 51, 52, 5, 42, 0, 0, 52, 53, 5, 1, 0, 0, 53,
		64, 1, 0, 0, 0, 54, 57, 3, 4, 2, 0, 55, 57, 3, 8, 4, 0, 56, 54, 1, 0, 0,
		0, 56, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 5, 42, 0, 0, 59, 60,
		7, 0, 0, 0, 60, 61, 5, 42, 0, 0, 61, 62, 3, 16, 8, 0, 62, 64, 1, 0, 0,
		0, 63, 33, 1, 0, 0, 0, 63, 50, 1, 0, 0, 0, 63, 56, 1, 0, 0, 0, 64, 72,
		1, 0, 0, 0, 65, 66, 10, 3, 0, 0, 66, 67, 5, 42, 0, 0, 67, 68, 5, 17, 0,
		0, 68, 69, 5, 42, 0, 0, 69, 71, 3, 2, 1, 4, 70, 65, 1, 0, 0, 0, 71, 74,
		1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 3, 1, 0, 0, 0,
		74, 72, 1, 0, 0, 0, 75, 77, 5, 32, 0, 0, 76, 78, 3, 12, 6, 0, 77, 76, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 81, 3, 8, 4, 0, 80,
		75, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 5, 1, 0, 0, 0, 82, 83, 7, 1, 0,
		0, 83, 7, 1, 0, 0, 0, 84, 85, 5, 32, 0, 0, 85, 87, 5, 37, 0, 0, 86, 88,
		3, 10, 5, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0,
		89, 90, 5, 38, 0, 0, 90, 9, 1, 0, 0, 0, 91, 96, 3, 16, 8, 0, 92, 93, 5,
		41, 0, 0, 93, 95, 3, 16, 8, 0, 94, 92, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 11, 1, 0, 0, 0, 98, 96, 1,
		0, 0, 0, 99, 100, 5, 12, 0, 0, 100, 101, 3, 4, 2, 0, 101, 13, 1, 0, 0,
		0, 102, 104, 3, 6, 3, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 121, 5, 34, 0, 0, 106, 108, 3, 6, 3, 0, 107, 106,
		1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 121, 5, 35,
		0, 0, 110, 112, 3, 6, 3, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0,
		112, 114, 1, 0, 0, 0, 113, 115, 5, 13, 0, 0, 114, 113, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 5, 36, 0, 0, 117, 119,
		5, 39, 0, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 1, 0,
		0, 0, 120, 103, 1, 0, 0, 0, 120, 107, 1, 0, 0, 0, 120, 111, 1, 0, 0, 0,
		121, 15, 1, 0, 0, 0, 122, 130, 3, 14, 7, 0, 123, 130, 5, 18, 0, 0, 124,
		130, 5, 19, 0, 0, 125, 130, 5, 33, 0, 0, 126, 130, 3, 26, 13, 0, 127, 130,
		3, 22, 11, 0, 128, 130, 3, 18, 9, 0, 129, 122, 1, 0, 0, 0, 129, 123, 1,
		0, 0, 0, 129, 124, 1, 0, 0, 0, 129, 125, 1, 0, 0, 0, 129, 126, 1, 0, 0,
		0, 129, 127, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 17, 1, 0, 0, 0, 131,
		132, 5, 14, 0, 0, 132, 133, 3, 20, 10, 0, 133, 19, 1, 0, 0, 0, 134, 135,
		5, 34, 0, 0, 135, 136, 5, 41, 0, 0, 136, 140, 3, 20, 10, 0, 137, 138, 5,
		34, 0, 0, 138, 140, 5, 15, 0, 0, 139, 134, 1, 0, 0, 0, 139, 137, 1, 0,
		0, 0, 140, 21, 1, 0, 0, 0, 141, 142, 5, 14, 0, 0, 142, 143, 3, 24, 12,
		0, 143, 23, 1, 0, 0, 0, 144, 145, 5, 35, 0, 0, 145, 146, 5, 41, 0, 0, 146,
		150, 3, 24, 12, 0, 147, 148, 5, 35, 0, 0, 148, 150, 5, 15, 0, 0, 149, 144,
		1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 25, 1, 0, 0, 0, 151, 152, 5, 14,
		0, 0, 152, 153, 3, 28, 14, 0, 153, 27, 1, 0, 0, 0, 154, 155, 5, 36, 0,
		0, 155, 156, 5, 41, 0, 0, 156, 160, 3, 28, 14, 0, 157, 158, 5, 36, 0, 0,
		158, 160, 5, 15, 0, 0, 159, 154, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160,
		29, 1, 0, 0, 0, 21, 35, 38, 42, 46, 56, 63, 72, 77, 80, 87, 96, 103, 107,
		111, 114, 118, 120, 129, 139, 149, 159,
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

// JsonQueryParserInit initializes any static state used to implement JsonQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonQueryParserInit() {
	staticData := &JsonQueryParserStaticData
	staticData.once.Do(jsonqueryParserInit)
}

// NewJsonQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonQueryParser(input antlr.TokenStream) *JsonQueryParser {
	JsonQueryParserInit()
	this := new(JsonQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JsonQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JsonQuery.g4"

	return this
}

// JsonQueryParser tokens.
const (
	JsonQueryParserEOF              = antlr.TokenEOF
	JsonQueryParserT__0             = 1
	JsonQueryParserT__1             = 2
	JsonQueryParserT__2             = 3
	JsonQueryParserT__3             = 4
	JsonQueryParserT__4             = 5
	JsonQueryParserT__5             = 6
	JsonQueryParserT__6             = 7
	JsonQueryParserT__7             = 8
	JsonQueryParserT__8             = 9
	JsonQueryParserT__9             = 10
	JsonQueryParserT__10            = 11
	JsonQueryParserT__11            = 12
	JsonQueryParserT__12            = 13
	JsonQueryParserT__13            = 14
	JsonQueryParserT__14            = 15
	JsonQueryParserNOT              = 16
	JsonQueryParserLOGICAL_OPERATOR = 17
	JsonQueryParserBOOLEAN          = 18
	JsonQueryParserNULL             = 19
	JsonQueryParserIN               = 20
	JsonQueryParserEQ               = 21
	JsonQueryParserNE               = 22
	JsonQueryParserGT               = 23
	JsonQueryParserLT               = 24
	JsonQueryParserGE               = 25
	JsonQueryParserLE               = 26
	JsonQueryParserCO               = 27
	JsonQueryParserSW               = 28
	JsonQueryParserEW               = 29
	JsonQueryParserDGT              = 30
	JsonQueryParserDEQ              = 31
	JsonQueryParserATTRNAME         = 32
	JsonQueryParserVERSION          = 33
	JsonQueryParserSTRING           = 34
	JsonQueryParserDOUBLE           = 35
	JsonQueryParserINT              = 36
	JsonQueryParserLPAREN           = 37
	JsonQueryParserRPAREN           = 38
	JsonQueryParserEXP              = 39
	JsonQueryParserNEWLINE          = 40
	JsonQueryParserCOMMA            = 41
	JsonQueryParserSP               = 42
)

// JsonQueryParser rules.
const (
	JsonQueryParserRULE_root             = 0
	JsonQueryParserRULE_query            = 1
	JsonQueryParserRULE_attrPath         = 2
	JsonQueryParserRULE_typeAnnotation   = 3
	JsonQueryParserRULE_functionCall     = 4
	JsonQueryParserRULE_argList          = 5
	JsonQueryParserRULE_subAttr          = 6
	JsonQueryParserRULE_typedValue       = 7
	JsonQueryParserRULE_value            = 8
	JsonQueryParserRULE_listStrings      = 9
	JsonQueryParserRULE_subListOfStrings = 10
	JsonQueryParserRULE_listDoubles      = 11
	JsonQueryParserRULE_subListOfDoubles = 12
	JsonQueryParserRULE_listInts         = 13
	JsonQueryParserRULE_subListOfInts    = 14
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Query() IQueryContext
	EOF() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonQueryParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.query(0)
	}
	{
		p.SetState(31)
		p.Match(JsonQueryParserEOF)
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

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyAll(ctx *QueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	QueryContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *CompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *CompareExpContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *CompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *CompareExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *CompareExpContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *CompareExpContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *CompareExpContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *CompareExpContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *CompareExpContext) CO() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCO, 0)
}

func (s *CompareExpContext) SW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSW, 0)
}

func (s *CompareExpContext) EW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEW, 0)
}

func (s *CompareExpContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIN, 0)
}

func (s *CompareExpContext) DGT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDGT, 0)
}

func (s *CompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterCompareExp(s)
	}
}

func (s *CompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitCompareExp(s)
	}
}

func (s *CompareExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitCompareExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	QueryContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLPAREN, 0)
}

func (s *ParenExpContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *ParenExpContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserRPAREN, 0)
}

func (s *ParenExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNOT, 0)
}

func (s *ParenExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *ParenExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *ParenExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterParenExp(s)
	}
}

func (s *ParenExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitParenExp(s)
	}
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PresentExpContext struct {
	QueryContext
}

func NewPresentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PresentExpContext {
	var p = new(PresentExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *PresentExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresentExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *PresentExpContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
}

func (s *PresentExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterPresentExp(s)
	}
}

func (s *PresentExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitPresentExp(s)
	}
}

func (s *PresentExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitPresentExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExpContext struct {
	QueryContext
}

func NewLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpContext {
	var p = new(LogicalExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *LogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
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

	return t.(IQueryContext)
}

func (s *LogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *LogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *LogicalExpContext) LOGICAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLOGICAL_OPERATOR, 0)
}

func (s *LogicalExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterLogicalExp(s)
	}
}

func (s *LogicalExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitLogicalExp(s)
	}
}

func (s *LogicalExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLogicalExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *JsonQueryParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, JsonQueryParserRULE_query, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserNOT {
			{
				p.SetState(34)
				p.Match(JsonQueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(37)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(40)
			p.Match(JsonQueryParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(41)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(44)
			p.query(0)
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(45)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(48)
			p.Match(JsonQueryParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewPresentExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.AttrPath()
		}
		{
			p.SetState(51)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(54)
				p.AttrPath()
			}

		case 2:
			{
				p.SetState(55)
				p.FunctionCall()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(58)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2146435072) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(60)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_query)
			p.SetState(65)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(66)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(67)
				p.Match(JsonQueryParserLOGICAL_OPERATOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(68)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(69)
				p.query(4)
			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRNAME() antlr.TerminalNode
	SubAttr() ISubAttrContext
	FunctionCall() IFunctionCallContext

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
	return p
}

func InitEmptyAttrPathContext(p *AttrPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterAttrPath(s)
	}
}

func (s *AttrPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitAttrPath(s)
	}
}

func (s *AttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonQueryParserRULE_attrPath)
	var _la int

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(JsonQueryParserATTRNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserT__11 {
			{
				p.SetState(76)
				p.SubAttr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.FunctionCall()
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

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_typeAnnotation
	return p
}

func InitEmptyTypeAnnotationContext(p *TypeAnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_typeAnnotation
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitTypeAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonQueryParserRULE_typeAnnotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4092) != 0) {
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

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRNAME() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgList() IArgListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserATTRNAME, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserRPAREN, 0)
}

func (s *FunctionCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonQueryParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(JsonQueryParserATTRNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(JsonQueryParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&128849833980) != 0 {
		{
			p.SetState(86)
			p.ArgList()
		}

	}
	{
		p.SetState(89)
		p.Match(JsonQueryParserRPAREN)
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

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonQueryParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Value()
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JsonQueryParserCOMMA {
		{
			p.SetState(92)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Value()
		}

		p.SetState(98)
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

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrPath() IAttrPathContext

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
	return p
}

func InitEmptySubAttrContext(p *SubAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterSubAttr(s)
	}
}

func (s *SubAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitSubAttr(s)
	}
}

func (s *SubAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonQueryParserRULE_subAttr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(JsonQueryParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.AttrPath()
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

// ITypedValueContext is an interface to support dynamic dispatch.
type ITypedValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypedValueContext differentiates from other interfaces.
	IsTypedValueContext()
}

type TypedValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedValueContext() *TypedValueContext {
	var p = new(TypedValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_typedValue
	return p
}

func InitEmptyTypedValueContext(p *TypedValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_typedValue
}

func (*TypedValueContext) IsTypedValueContext() {}

func NewTypedValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedValueContext {
	var p = new(TypedValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_typedValue

	return p
}

func (s *TypedValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedValueContext) CopyAll(ctx *TypedValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypedValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedDoubleContext struct {
	TypedValueContext
}

func NewTypedDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDoubleContext {
	var p = new(TypedDoubleContext)

	InitEmptyTypedValueContext(&p.TypedValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypedValueContext))

	return p
}

func (s *TypedDoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *TypedDoubleContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *TypedDoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterTypedDouble(s)
	}
}

func (s *TypedDoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitTypedDouble(s)
	}
}

func (s *TypedDoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitTypedDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedStringContext struct {
	TypedValueContext
}

func NewTypedStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedStringContext {
	var p = new(TypedStringContext)

	InitEmptyTypedValueContext(&p.TypedValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypedValueContext))

	return p
}

func (s *TypedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *TypedStringContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *TypedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterTypedString(s)
	}
}

func (s *TypedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitTypedString(s)
	}
}

func (s *TypedStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitTypedString(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedIntegerContext struct {
	TypedValueContext
}

func NewTypedIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedIntegerContext {
	var p = new(TypedIntegerContext)

	InitEmptyTypedValueContext(&p.TypedValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypedValueContext))

	return p
}

func (s *TypedIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedIntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *TypedIntegerContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *TypedIntegerContext) EXP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEXP, 0)
}

func (s *TypedIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterTypedInteger(s)
	}
}

func (s *TypedIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitTypedInteger(s)
	}
}

func (s *TypedIntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitTypedInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) TypedValue() (localctx ITypedValueContext) {
	localctx = NewTypedValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonQueryParserRULE_typedValue)
	var _la int

	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypedStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4092) != 0 {
			{
				p.SetState(102)
				p.TypeAnnotation()
			}

		}
		{
			p.SetState(105)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTypedDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4092) != 0 {
			{
				p.SetState(106)
				p.TypeAnnotation()
			}

		}
		{
			p.SetState(109)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTypedIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4092) != 0 {
			{
				p.SetState(110)
				p.TypeAnnotation()
			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserT__12 {
			{
				p.SetState(113)
				p.Match(JsonQueryParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(116)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(117)
				p.Match(JsonQueryParserEXP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedValContext struct {
	ValueContext
}

func NewTypedValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedValContext {
	var p = new(TypedValContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *TypedValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedValContext) TypedValue() ITypedValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedValueContext)
}

func (s *TypedValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterTypedVal(s)
	}
}

func (s *TypedValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitTypedVal(s)
	}
}

func (s *TypedValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitTypedVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfDoublesContext struct {
	ValueContext
}

func NewListOfDoublesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfDoublesContext {
	var p = new(ListOfDoublesContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfDoublesContext) ListDoubles() IListDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListDoublesContext)
}

func (s *ListOfDoublesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterListOfDoubles(s)
	}
}

func (s *ListOfDoublesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitListOfDoubles(s)
	}
}

func (s *ListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfStringsContext struct {
	ValueContext
}

func NewListOfStringsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfStringsContext {
	var p = new(ListOfStringsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfStringsContext) ListStrings() IListStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *ListOfStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterListOfStrings(s)
	}
}

func (s *ListOfStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitListOfStrings(s)
	}
}

func (s *ListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserBOOLEAN, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	ValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNULL, 0)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type VersionContext struct {
	ValueContext
}

func NewVersionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VersionContext {
	var p = new(VersionContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) VERSION() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserVERSION, 0)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfIntsContext struct {
	ValueContext
}

func NewListOfIntsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfIntsContext {
	var p = new(ListOfIntsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfIntsContext) ListInts() IListIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListIntsContext)
}

func (s *ListOfIntsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterListOfInts(s)
	}
}

func (s *ListOfIntsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitListOfInts(s)
	}
}

func (s *ListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonQueryParserRULE_value)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypedValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.TypedValue()
		}

	case 2:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(JsonQueryParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewVersionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
			p.Match(JsonQueryParserVERSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewListOfIntsContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.ListInts()
		}

	case 6:
		localctx = NewListOfDoublesContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)
			p.ListDoubles()
		}

	case 7:
		localctx = NewListOfStringsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(128)
			p.ListStrings()
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

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfStrings() ISubListOfStringsContext

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
	return p
}

func InitEmptyListStringsContext(p *ListStringsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterListStrings(s)
	}
}

func (s *ListStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitListStrings(s)
	}
}

func (s *ListStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonQueryParserRULE_listStrings)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(JsonQueryParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.SubListOfStrings()
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

// ISubListOfStringsContext is an interface to support dynamic dispatch.
type ISubListOfStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfStrings() ISubListOfStringsContext

	// IsSubListOfStringsContext differentiates from other interfaces.
	IsSubListOfStringsContext()
}

type SubListOfStringsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfStringsContext() *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings
	return p
}

func InitEmptySubListOfStringsContext(p *SubListOfStringsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings
}

func (*SubListOfStringsContext) IsSubListOfStringsContext() {}

func NewSubListOfStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings

	return p
}

func (s *SubListOfStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfStringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *SubListOfStringsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *SubListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterSubListOfStrings(s)
	}
}

func (s *SubListOfStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitSubListOfStrings(s)
	}
}

func (s *SubListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfStrings() (localctx ISubListOfStringsContext) {
	localctx = NewSubListOfStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonQueryParserRULE_subListOfStrings)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.SubListOfStrings()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(JsonQueryParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IListDoublesContext is an interface to support dynamic dispatch.
type IListDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfDoubles() ISubListOfDoublesContext

	// IsListDoublesContext differentiates from other interfaces.
	IsListDoublesContext()
}

type ListDoublesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListDoublesContext() *ListDoublesContext {
	var p = new(ListDoublesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
	return p
}

func InitEmptyListDoublesContext(p *ListDoublesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
}

func (*ListDoublesContext) IsListDoublesContext() {}

func NewListDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDoublesContext {
	var p = new(ListDoublesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listDoubles

	return p
}

func (s *ListDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *ListDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDoublesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterListDoubles(s)
	}
}

func (s *ListDoublesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitListDoubles(s)
	}
}

func (s *ListDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListDoubles() (localctx IListDoublesContext) {
	localctx = NewListDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonQueryParserRULE_listDoubles)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(JsonQueryParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.SubListOfDoubles()
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

// ISubListOfDoublesContext is an interface to support dynamic dispatch.
type ISubListOfDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfDoubles() ISubListOfDoublesContext

	// IsSubListOfDoublesContext differentiates from other interfaces.
	IsSubListOfDoublesContext()
}

type SubListOfDoublesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfDoublesContext() *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles
	return p
}

func InitEmptySubListOfDoublesContext(p *SubListOfDoublesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles
}

func (*SubListOfDoublesContext) IsSubListOfDoublesContext() {}

func NewSubListOfDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles

	return p
}

func (s *SubListOfDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfDoublesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *SubListOfDoublesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *SubListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfDoublesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterSubListOfDoubles(s)
	}
}

func (s *SubListOfDoublesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitSubListOfDoubles(s)
	}
}

func (s *SubListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfDoubles() (localctx ISubListOfDoublesContext) {
	localctx = NewSubListOfDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonQueryParserRULE_subListOfDoubles)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.SubListOfDoubles()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(JsonQueryParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IListIntsContext is an interface to support dynamic dispatch.
type IListIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfInts() ISubListOfIntsContext

	// IsListIntsContext differentiates from other interfaces.
	IsListIntsContext()
}

type ListIntsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIntsContext() *ListIntsContext {
	var p = new(ListIntsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
	return p
}

func InitEmptyListIntsContext(p *ListIntsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
}

func (*ListIntsContext) IsListIntsContext() {}

func NewListIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIntsContext {
	var p = new(ListIntsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listInts

	return p
}

func (s *ListIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *ListIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIntsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterListInts(s)
	}
}

func (s *ListIntsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitListInts(s)
	}
}

func (s *ListIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListInts() (localctx IListIntsContext) {
	localctx = NewListIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JsonQueryParserRULE_listInts)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(JsonQueryParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.SubListOfInts()
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

// ISubListOfIntsContext is an interface to support dynamic dispatch.
type ISubListOfIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfInts() ISubListOfIntsContext

	// IsSubListOfIntsContext differentiates from other interfaces.
	IsSubListOfIntsContext()
}

type SubListOfIntsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfIntsContext() *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfInts
	return p
}

func InitEmptySubListOfIntsContext(p *SubListOfIntsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfInts
}

func (*SubListOfIntsContext) IsSubListOfIntsContext() {}

func NewSubListOfIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfInts

	return p
}

func (s *SubListOfIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfIntsContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *SubListOfIntsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *SubListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfIntsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.EnterSubListOfInts(s)
	}
}

func (s *SubListOfIntsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonQueryListener); ok {
		listenerT.ExitSubListOfInts(s)
	}
}

func (s *SubListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfInts() (localctx ISubListOfIntsContext) {
	localctx = NewSubListOfIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JsonQueryParserRULE_subListOfInts)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.SubListOfInts()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(JsonQueryParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (p *JsonQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonQueryParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
