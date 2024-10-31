// Code generated from Plan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package planparserv2

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

type PlanLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PlanLexerLexerStaticData struct {
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

func planlexerLexerInit() {
	staticData := &PlanLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'['", "','", "']'", "'{'", "'}'", "'<'", "'<='",
		"'>'", "'>='", "'=='", "'!='", "", "", "", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'**'", "'<<'", "'>>'", "'&'", "'|'", "'^'", "", "", "'~'", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "'$meta'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "LBRACE", "RBRACE", "LT", "LE", "GT", "GE",
		"EQ", "NE", "LIKE", "EXISTS", "TEXTMATCH", "ADD", "SUB", "MUL", "DIV",
		"MOD", "POW", "SHL", "SHR", "BAND", "BOR", "BXOR", "AND", "OR", "BNOT",
		"NOT", "IN", "EmptyArray", "JSONContains", "JSONContainsAll", "JSONContainsAny",
		"ArrayContains", "ArrayContainsAll", "ArrayContainsAny", "ArrayLength",
		"BooleanConstant", "IntegerConstant", "FloatingConstant", "Identifier",
		"Meta", "StringLiteral", "JSONIdentifier", "Whitespace", "Newline",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "LBRACE", "RBRACE", "LT", "LE",
		"GT", "GE", "EQ", "NE", "LIKE", "EXISTS", "TEXTMATCH", "ADD", "SUB",
		"MUL", "DIV", "MOD", "POW", "SHL", "SHR", "BAND", "BOR", "BXOR", "AND",
		"OR", "BNOT", "NOT", "IN", "EmptyArray", "JSONContains", "JSONContainsAll",
		"JSONContainsAny", "ArrayContains", "ArrayContainsAll", "ArrayContainsAny",
		"ArrayLength", "BooleanConstant", "IntegerConstant", "FloatingConstant",
		"Identifier", "Meta", "StringLiteral", "JSONIdentifier", "EncodingPrefix",
		"DoubleSCharSequence", "SingleSCharSequence", "DoubleSChar", "SingleSChar",
		"Nondigit", "Digit", "BinaryConstant", "DecimalConstant", "OctalConstant",
		"HexadecimalConstant", "NonzeroDigit", "OctalDigit", "HexadecimalDigit",
		"HexQuad", "UniversalCharacterName", "DecimalFloatingConstant", "HexadecimalFloatingConstant",
		"FractionalConstant", "ExponentPart", "DigitSequence", "HexadecimalFractionalConstant",
		"HexadecimalDigitSequence", "BinaryExponentPart", "EscapeSequence",
		"Whitespace", "Newline",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 49, 789, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 188, 8, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3,
		14, 202, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 231,
		8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3,
		27, 263, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 269, 8, 28, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 277, 8, 30, 1, 31, 1, 31, 1, 31,
		1, 31, 3, 31, 283, 8, 31, 1, 32, 1, 32, 1, 32, 5, 32, 288, 8, 32, 10, 32,
		12, 32, 291, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3,
		33, 321, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 357, 8, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		3, 35, 393, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3,
		36, 423, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 461, 8, 37, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 499, 8, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3,
		39, 525, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 554,
		8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 560, 8, 41, 1, 42, 1, 42, 3,
		42, 564, 8, 42, 1, 43, 1, 43, 1, 43, 5, 43, 569, 8, 43, 10, 43, 12, 43,
		572, 9, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 3, 45, 581,
		8, 45, 1, 45, 1, 45, 3, 45, 585, 8, 45, 1, 45, 1, 45, 1, 45, 3, 45, 590,
		8, 45, 1, 45, 3, 45, 593, 8, 45, 1, 46, 1, 46, 3, 46, 597, 8, 46, 1, 46,
		1, 46, 1, 46, 3, 46, 602, 8, 46, 1, 46, 1, 46, 4, 46, 606, 8, 46, 11, 46,
		12, 46, 607, 1, 47, 1, 47, 1, 47, 3, 47, 613, 8, 47, 1, 48, 4, 48, 616,
		8, 48, 11, 48, 12, 48, 617, 1, 49, 4, 49, 621, 8, 49, 11, 49, 12, 49, 622,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 632, 8, 50, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 641, 8, 51, 1, 52,
		1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 4, 54, 650, 8, 54, 11, 54, 12,
		54, 651, 1, 55, 1, 55, 5, 55, 656, 8, 55, 10, 55, 12, 55, 659, 9, 55, 1,
		55, 3, 55, 662, 8, 55, 1, 56, 1, 56, 5, 56, 666, 8, 56, 10, 56, 12, 56,
		669, 9, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1,
		60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62,
		1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 3, 62, 696, 8, 62, 1, 63, 1,
		63, 3, 63, 700, 8, 63, 1, 63, 1, 63, 1, 63, 3, 63, 705, 8, 63, 1, 64, 1,
		64, 1, 64, 1, 64, 3, 64, 711, 8, 64, 1, 64, 1, 64, 1, 65, 3, 65, 716, 8,
		65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 3, 65, 723, 8, 65, 1, 66, 1, 66,
		3, 66, 727, 8, 66, 1, 66, 1, 66, 1, 67, 4, 67, 732, 8, 67, 11, 67, 12,
		67, 733, 1, 68, 3, 68, 737, 8, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 3,
		68, 744, 8, 68, 1, 69, 4, 69, 747, 8, 69, 11, 69, 12, 69, 748, 1, 70, 1,
		70, 3, 70, 753, 8, 70, 1, 70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71,
		3, 71, 762, 8, 71, 1, 71, 3, 71, 765, 8, 71, 1, 71, 1, 71, 1, 71, 1, 71,
		1, 71, 3, 71, 772, 8, 71, 1, 72, 4, 72, 775, 8, 72, 11, 72, 12, 72, 776,
		1, 72, 1, 72, 1, 73, 1, 73, 3, 73, 783, 8, 73, 1, 73, 3, 73, 786, 8, 73,
		1, 73, 1, 73, 0, 0, 74, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 0, 97, 0, 99, 0, 101, 0, 103, 0, 105, 0, 107,
		0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119, 0, 121, 0, 123, 0, 125,
		0, 127, 0, 129, 0, 131, 0, 133, 0, 135, 0, 137, 0, 139, 0, 141, 0, 143,
		0, 145, 48, 147, 49, 1, 0, 16, 3, 0, 76, 76, 85, 85, 117, 117, 4, 0, 10,
		10, 13, 13, 34, 34, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 3, 0,
		65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 66, 66, 98, 98, 1, 0, 48,
		49, 2, 0, 88, 88, 120, 120, 1, 0, 49, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65,
		70, 97, 102, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 80, 80,
		112, 112, 10, 0, 34, 34, 39, 39, 63, 63, 92, 92, 97, 98, 102, 102, 110,
		110, 114, 114, 116, 116, 118, 118, 2, 0, 9, 9, 32, 32, 831, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0,
		0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0,
		0, 0, 145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 1, 149, 1, 0, 0, 0, 3, 151,
		1, 0, 0, 0, 5, 153, 1, 0, 0, 0, 7, 155, 1, 0, 0, 0, 9, 157, 1, 0, 0, 0,
		11, 159, 1, 0, 0, 0, 13, 161, 1, 0, 0, 0, 15, 163, 1, 0, 0, 0, 17, 165,
		1, 0, 0, 0, 19, 168, 1, 0, 0, 0, 21, 170, 1, 0, 0, 0, 23, 173, 1, 0, 0,
		0, 25, 176, 1, 0, 0, 0, 27, 187, 1, 0, 0, 0, 29, 201, 1, 0, 0, 0, 31, 230,
		1, 0, 0, 0, 33, 232, 1, 0, 0, 0, 35, 234, 1, 0, 0, 0, 37, 236, 1, 0, 0,
		0, 39, 238, 1, 0, 0, 0, 41, 240, 1, 0, 0, 0, 43, 242, 1, 0, 0, 0, 45, 245,
		1, 0, 0, 0, 47, 248, 1, 0, 0, 0, 49, 251, 1, 0, 0, 0, 51, 253, 1, 0, 0,
		0, 53, 255, 1, 0, 0, 0, 55, 262, 1, 0, 0, 0, 57, 268, 1, 0, 0, 0, 59, 270,
		1, 0, 0, 0, 61, 276, 1, 0, 0, 0, 63, 282, 1, 0, 0, 0, 65, 284, 1, 0, 0,
		0, 67, 320, 1, 0, 0, 0, 69, 356, 1, 0, 0, 0, 71, 392, 1, 0, 0, 0, 73, 422,
		1, 0, 0, 0, 75, 460, 1, 0, 0, 0, 77, 498, 1, 0, 0, 0, 79, 524, 1, 0, 0,
		0, 81, 553, 1, 0, 0, 0, 83, 559, 1, 0, 0, 0, 85, 563, 1, 0, 0, 0, 87, 565,
		1, 0, 0, 0, 89, 573, 1, 0, 0, 0, 91, 580, 1, 0, 0, 0, 93, 596, 1, 0, 0,
		0, 95, 612, 1, 0, 0, 0, 97, 615, 1, 0, 0, 0, 99, 620, 1, 0, 0, 0, 101,
		631, 1, 0, 0, 0, 103, 640, 1, 0, 0, 0, 105, 642, 1, 0, 0, 0, 107, 644,
		1, 0, 0, 0, 109, 646, 1, 0, 0, 0, 111, 661, 1, 0, 0, 0, 113, 663, 1, 0,
		0, 0, 115, 670, 1, 0, 0, 0, 117, 674, 1, 0, 0, 0, 119, 676, 1, 0, 0, 0,
		121, 678, 1, 0, 0, 0, 123, 680, 1, 0, 0, 0, 125, 695, 1, 0, 0, 0, 127,
		704, 1, 0, 0, 0, 129, 706, 1, 0, 0, 0, 131, 722, 1, 0, 0, 0, 133, 724,
		1, 0, 0, 0, 135, 731, 1, 0, 0, 0, 137, 743, 1, 0, 0, 0, 139, 746, 1, 0,
		0, 0, 141, 750, 1, 0, 0, 0, 143, 771, 1, 0, 0, 0, 145, 774, 1, 0, 0, 0,
		147, 785, 1, 0, 0, 0, 149, 150, 5, 40, 0, 0, 150, 2, 1, 0, 0, 0, 151, 152,
		5, 41, 0, 0, 152, 4, 1, 0, 0, 0, 153, 154, 5, 91, 0, 0, 154, 6, 1, 0, 0,
		0, 155, 156, 5, 44, 0, 0, 156, 8, 1, 0, 0, 0, 157, 158, 5, 93, 0, 0, 158,
		10, 1, 0, 0, 0, 159, 160, 5, 123, 0, 0, 160, 12, 1, 0, 0, 0, 161, 162,
		5, 125, 0, 0, 162, 14, 1, 0, 0, 0, 163, 164, 5, 60, 0, 0, 164, 16, 1, 0,
		0, 0, 165, 166, 5, 60, 0, 0, 166, 167, 5, 61, 0, 0, 167, 18, 1, 0, 0, 0,
		168, 169, 5, 62, 0, 0, 169, 20, 1, 0, 0, 0, 170, 171, 5, 62, 0, 0, 171,
		172, 5, 61, 0, 0, 172, 22, 1, 0, 0, 0, 173, 174, 5, 61, 0, 0, 174, 175,
		5, 61, 0, 0, 175, 24, 1, 0, 0, 0, 176, 177, 5, 33, 0, 0, 177, 178, 5, 61,
		0, 0, 178, 26, 1, 0, 0, 0, 179, 180, 5, 108, 0, 0, 180, 181, 5, 105, 0,
		0, 181, 182, 5, 107, 0, 0, 182, 188, 5, 101, 0, 0, 183, 184, 5, 76, 0,
		0, 184, 185, 5, 73, 0, 0, 185, 186, 5, 75, 0, 0, 186, 188, 5, 69, 0, 0,
		187, 179, 1, 0, 0, 0, 187, 183, 1, 0, 0, 0, 188, 28, 1, 0, 0, 0, 189, 190,
		5, 101, 0, 0, 190, 191, 5, 120, 0, 0, 191, 192, 5, 105, 0, 0, 192, 193,
		5, 115, 0, 0, 193, 194, 5, 116, 0, 0, 194, 202, 5, 115, 0, 0, 195, 196,
		5, 69, 0, 0, 196, 197, 5, 88, 0, 0, 197, 198, 5, 73, 0, 0, 198, 199, 5,
		83, 0, 0, 199, 200, 5, 84, 0, 0, 200, 202, 5, 83, 0, 0, 201, 189, 1, 0,
		0, 0, 201, 195, 1, 0, 0, 0, 202, 30, 1, 0, 0, 0, 203, 204, 5, 84, 0, 0,
		204, 205, 5, 101, 0, 0, 205, 206, 5, 120, 0, 0, 206, 207, 5, 116, 0, 0,
		207, 208, 5, 77, 0, 0, 208, 209, 5, 97, 0, 0, 209, 210, 5, 116, 0, 0, 210,
		211, 5, 99, 0, 0, 211, 231, 5, 104, 0, 0, 212, 213, 5, 116, 0, 0, 213,
		214, 5, 101, 0, 0, 214, 215, 5, 120, 0, 0, 215, 216, 5, 116, 0, 0, 216,
		217, 5, 109, 0, 0, 217, 218, 5, 97, 0, 0, 218, 219, 5, 116, 0, 0, 219,
		220, 5, 99, 0, 0, 220, 231, 5, 104, 0, 0, 221, 222, 5, 84, 0, 0, 222, 223,
		5, 69, 0, 0, 223, 224, 5, 88, 0, 0, 224, 225, 5, 84, 0, 0, 225, 226, 5,
		77, 0, 0, 226, 227, 5, 65, 0, 0, 227, 228, 5, 84, 0, 0, 228, 229, 5, 67,
		0, 0, 229, 231, 5, 72, 0, 0, 230, 203, 1, 0, 0, 0, 230, 212, 1, 0, 0, 0,
		230, 221, 1, 0, 0, 0, 231, 32, 1, 0, 0, 0, 232, 233, 5, 43, 0, 0, 233,
		34, 1, 0, 0, 0, 234, 235, 5, 45, 0, 0, 235, 36, 1, 0, 0, 0, 236, 237, 5,
		42, 0, 0, 237, 38, 1, 0, 0, 0, 238, 239, 5, 47, 0, 0, 239, 40, 1, 0, 0,
		0, 240, 241, 5, 37, 0, 0, 241, 42, 1, 0, 0, 0, 242, 243, 5, 42, 0, 0, 243,
		244, 5, 42, 0, 0, 244, 44, 1, 0, 0, 0, 245, 246, 5, 60, 0, 0, 246, 247,
		5, 60, 0, 0, 247, 46, 1, 0, 0, 0, 248, 249, 5, 62, 0, 0, 249, 250, 5, 62,
		0, 0, 250, 48, 1, 0, 0, 0, 251, 252, 5, 38, 0, 0, 252, 50, 1, 0, 0, 0,
		253, 254, 5, 124, 0, 0, 254, 52, 1, 0, 0, 0, 255, 256, 5, 94, 0, 0, 256,
		54, 1, 0, 0, 0, 257, 258, 5, 38, 0, 0, 258, 263, 5, 38, 0, 0, 259, 260,
		5, 97, 0, 0, 260, 261, 5, 110, 0, 0, 261, 263, 5, 100, 0, 0, 262, 257,
		1, 0, 0, 0, 262, 259, 1, 0, 0, 0, 263, 56, 1, 0, 0, 0, 264, 265, 5, 124,
		0, 0, 265, 269, 5, 124, 0, 0, 266, 267, 5, 111, 0, 0, 267, 269, 5, 114,
		0, 0, 268, 264, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 58, 1, 0, 0, 0,
		270, 271, 5, 126, 0, 0, 271, 60, 1, 0, 0, 0, 272, 277, 5, 33, 0, 0, 273,
		274, 5, 110, 0, 0, 274, 275, 5, 111, 0, 0, 275, 277, 5, 116, 0, 0, 276,
		272, 1, 0, 0, 0, 276, 273, 1, 0, 0, 0, 277, 62, 1, 0, 0, 0, 278, 279, 5,
		105, 0, 0, 279, 283, 5, 110, 0, 0, 280, 281, 5, 73, 0, 0, 281, 283, 5,
		78, 0, 0, 282, 278, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 64, 1, 0, 0,
		0, 284, 289, 5, 91, 0, 0, 285, 288, 3, 145, 72, 0, 286, 288, 3, 147, 73,
		0, 287, 285, 1, 0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289,
		287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 292, 1, 0, 0, 0, 291, 289,
		1, 0, 0, 0, 292, 293, 5, 93, 0, 0, 293, 66, 1, 0, 0, 0, 294, 295, 5, 106,
		0, 0, 295, 296, 5, 115, 0, 0, 296, 297, 5, 111, 0, 0, 297, 298, 5, 110,
		0, 0, 298, 299, 5, 95, 0, 0, 299, 300, 5, 99, 0, 0, 300, 301, 5, 111, 0,
		0, 301, 302, 5, 110, 0, 0, 302, 303, 5, 116, 0, 0, 303, 304, 5, 97, 0,
		0, 304, 305, 5, 105, 0, 0, 305, 306, 5, 110, 0, 0, 306, 321, 5, 115, 0,
		0, 307, 308, 5, 74, 0, 0, 308, 309, 5, 83, 0, 0, 309, 310, 5, 79, 0, 0,
		310, 311, 5, 78, 0, 0, 311, 312, 5, 95, 0, 0, 312, 313, 5, 67, 0, 0, 313,
		314, 5, 79, 0, 0, 314, 315, 5, 78, 0, 0, 315, 316, 5, 84, 0, 0, 316, 317,
		5, 65, 0, 0, 317, 318, 5, 73, 0, 0, 318, 319, 5, 78, 0, 0, 319, 321, 5,
		83, 0, 0, 320, 294, 1, 0, 0, 0, 320, 307, 1, 0, 0, 0, 321, 68, 1, 0, 0,
		0, 322, 323, 5, 106, 0, 0, 323, 324, 5, 115, 0, 0, 324, 325, 5, 111, 0,
		0, 325, 326, 5, 110, 0, 0, 326, 327, 5, 95, 0, 0, 327, 328, 5, 99, 0, 0,
		328, 329, 5, 111, 0, 0, 329, 330, 5, 110, 0, 0, 330, 331, 5, 116, 0, 0,
		331, 332, 5, 97, 0, 0, 332, 333, 5, 105, 0, 0, 333, 334, 5, 110, 0, 0,
		334, 335, 5, 115, 0, 0, 335, 336, 5, 95, 0, 0, 336, 337, 5, 97, 0, 0, 337,
		338, 5, 108, 0, 0, 338, 357, 5, 108, 0, 0, 339, 340, 5, 74, 0, 0, 340,
		341, 5, 83, 0, 0, 341, 342, 5, 79, 0, 0, 342, 343, 5, 78, 0, 0, 343, 344,
		5, 95, 0, 0, 344, 345, 5, 67, 0, 0, 345, 346, 5, 79, 0, 0, 346, 347, 5,
		78, 0, 0, 347, 348, 5, 84, 0, 0, 348, 349, 5, 65, 0, 0, 349, 350, 5, 73,
		0, 0, 350, 351, 5, 78, 0, 0, 351, 352, 5, 83, 0, 0, 352, 353, 5, 95, 0,
		0, 353, 354, 5, 65, 0, 0, 354, 355, 5, 76, 0, 0, 355, 357, 5, 76, 0, 0,
		356, 322, 1, 0, 0, 0, 356, 339, 1, 0, 0, 0, 357, 70, 1, 0, 0, 0, 358, 359,
		5, 106, 0, 0, 359, 360, 5, 115, 0, 0, 360, 361, 5, 111, 0, 0, 361, 362,
		5, 110, 0, 0, 362, 363, 5, 95, 0, 0, 363, 364, 5, 99, 0, 0, 364, 365, 5,
		111, 0, 0, 365, 366, 5, 110, 0, 0, 366, 367, 5, 116, 0, 0, 367, 368, 5,
		97, 0, 0, 368, 369, 5, 105, 0, 0, 369, 370, 5, 110, 0, 0, 370, 371, 5,
		115, 0, 0, 371, 372, 5, 95, 0, 0, 372, 373, 5, 97, 0, 0, 373, 374, 5, 110,
		0, 0, 374, 393, 5, 121, 0, 0, 375, 376, 5, 74, 0, 0, 376, 377, 5, 83, 0,
		0, 377, 378, 5, 79, 0, 0, 378, 379, 5, 78, 0, 0, 379, 380, 5, 95, 0, 0,
		380, 381, 5, 67, 0, 0, 381, 382, 5, 79, 0, 0, 382, 383, 5, 78, 0, 0, 383,
		384, 5, 84, 0, 0, 384, 385, 5, 65, 0, 0, 385, 386, 5, 73, 0, 0, 386, 387,
		5, 78, 0, 0, 387, 388, 5, 83, 0, 0, 388, 389, 5, 95, 0, 0, 389, 390, 5,
		65, 0, 0, 390, 391, 5, 78, 0, 0, 391, 393, 5, 89, 0, 0, 392, 358, 1, 0,
		0, 0, 392, 375, 1, 0, 0, 0, 393, 72, 1, 0, 0, 0, 394, 395, 5, 97, 0, 0,
		395, 396, 5, 114, 0, 0, 396, 397, 5, 114, 0, 0, 397, 398, 5, 97, 0, 0,
		398, 399, 5, 121, 0, 0, 399, 400, 5, 95, 0, 0, 400, 401, 5, 99, 0, 0, 401,
		402, 5, 111, 0, 0, 402, 403, 5, 110, 0, 0, 403, 404, 5, 116, 0, 0, 404,
		405, 5, 97, 0, 0, 405, 406, 5, 105, 0, 0, 406, 407, 5, 110, 0, 0, 407,
		423, 5, 115, 0, 0, 408, 409, 5, 65, 0, 0, 409, 410, 5, 82, 0, 0, 410, 411,
		5, 82, 0, 0, 411, 412, 5, 65, 0, 0, 412, 413, 5, 89, 0, 0, 413, 414, 5,
		95, 0, 0, 414, 415, 5, 67, 0, 0, 415, 416, 5, 79, 0, 0, 416, 417, 5, 78,
		0, 0, 417, 418, 5, 84, 0, 0, 418, 419, 5, 65, 0, 0, 419, 420, 5, 73, 0,
		0, 420, 421, 5, 78, 0, 0, 421, 423, 5, 83, 0, 0, 422, 394, 1, 0, 0, 0,
		422, 408, 1, 0, 0, 0, 423, 74, 1, 0, 0, 0, 424, 425, 5, 97, 0, 0, 425,
		426, 5, 114, 0, 0, 426, 427, 5, 114, 0, 0, 427, 428, 5, 97, 0, 0, 428,
		429, 5, 121, 0, 0, 429, 430, 5, 95, 0, 0, 430, 431, 5, 99, 0, 0, 431, 432,
		5, 111, 0, 0, 432, 433, 5, 110, 0, 0, 433, 434, 5, 116, 0, 0, 434, 435,
		5, 97, 0, 0, 435, 436, 5, 105, 0, 0, 436, 437, 5, 110, 0, 0, 437, 438,
		5, 115, 0, 0, 438, 439, 5, 95, 0, 0, 439, 440, 5, 97, 0, 0, 440, 441, 5,
		108, 0, 0, 441, 461, 5, 108, 0, 0, 442, 443, 5, 65, 0, 0, 443, 444, 5,
		82, 0, 0, 444, 445, 5, 82, 0, 0, 445, 446, 5, 65, 0, 0, 446, 447, 5, 89,
		0, 0, 447, 448, 5, 95, 0, 0, 448, 449, 5, 67, 0, 0, 449, 450, 5, 79, 0,
		0, 450, 451, 5, 78, 0, 0, 451, 452, 5, 84, 0, 0, 452, 453, 5, 65, 0, 0,
		453, 454, 5, 73, 0, 0, 454, 455, 5, 78, 0, 0, 455, 456, 5, 83, 0, 0, 456,
		457, 5, 95, 0, 0, 457, 458, 5, 65, 0, 0, 458, 459, 5, 76, 0, 0, 459, 461,
		5, 76, 0, 0, 460, 424, 1, 0, 0, 0, 460, 442, 1, 0, 0, 0, 461, 76, 1, 0,
		0, 0, 462, 463, 5, 97, 0, 0, 463, 464, 5, 114, 0, 0, 464, 465, 5, 114,
		0, 0, 465, 466, 5, 97, 0, 0, 466, 467, 5, 121, 0, 0, 467, 468, 5, 95, 0,
		0, 468, 469, 5, 99, 0, 0, 469, 470, 5, 111, 0, 0, 470, 471, 5, 110, 0,
		0, 471, 472, 5, 116, 0, 0, 472, 473, 5, 97, 0, 0, 473, 474, 5, 105, 0,
		0, 474, 475, 5, 110, 0, 0, 475, 476, 5, 115, 0, 0, 476, 477, 5, 95, 0,
		0, 477, 478, 5, 97, 0, 0, 478, 479, 5, 110, 0, 0, 479, 499, 5, 121, 0,
		0, 480, 481, 5, 65, 0, 0, 481, 482, 5, 82, 0, 0, 482, 483, 5, 82, 0, 0,
		483, 484, 5, 65, 0, 0, 484, 485, 5, 89, 0, 0, 485, 486, 5, 95, 0, 0, 486,
		487, 5, 67, 0, 0, 487, 488, 5, 79, 0, 0, 488, 489, 5, 78, 0, 0, 489, 490,
		5, 84, 0, 0, 490, 491, 5, 65, 0, 0, 491, 492, 5, 73, 0, 0, 492, 493, 5,
		78, 0, 0, 493, 494, 5, 83, 0, 0, 494, 495, 5, 95, 0, 0, 495, 496, 5, 65,
		0, 0, 496, 497, 5, 78, 0, 0, 497, 499, 5, 89, 0, 0, 498, 462, 1, 0, 0,
		0, 498, 480, 1, 0, 0, 0, 499, 78, 1, 0, 0, 0, 500, 501, 5, 97, 0, 0, 501,
		502, 5, 114, 0, 0, 502, 503, 5, 114, 0, 0, 503, 504, 5, 97, 0, 0, 504,
		505, 5, 121, 0, 0, 505, 506, 5, 95, 0, 0, 506, 507, 5, 108, 0, 0, 507,
		508, 5, 101, 0, 0, 508, 509, 5, 110, 0, 0, 509, 510, 5, 103, 0, 0, 510,
		511, 5, 116, 0, 0, 511, 525, 5, 104, 0, 0, 512, 513, 5, 65, 0, 0, 513,
		514, 5, 82, 0, 0, 514, 515, 5, 82, 0, 0, 515, 516, 5, 65, 0, 0, 516, 517,
		5, 89, 0, 0, 517, 518, 5, 95, 0, 0, 518, 519, 5, 76, 0, 0, 519, 520, 5,
		69, 0, 0, 520, 521, 5, 78, 0, 0, 521, 522, 5, 71, 0, 0, 522, 523, 5, 84,
		0, 0, 523, 525, 5, 72, 0, 0, 524, 500, 1, 0, 0, 0, 524, 512, 1, 0, 0, 0,
		525, 80, 1, 0, 0, 0, 526, 527, 5, 116, 0, 0, 527, 528, 5, 114, 0, 0, 528,
		529, 5, 117, 0, 0, 529, 554, 5, 101, 0, 0, 530, 531, 5, 84, 0, 0, 531,
		532, 5, 114, 0, 0, 532, 533, 5, 117, 0, 0, 533, 554, 5, 101, 0, 0, 534,
		535, 5, 84, 0, 0, 535, 536, 5, 82, 0, 0, 536, 537, 5, 85, 0, 0, 537, 554,
		5, 69, 0, 0, 538, 539, 5, 102, 0, 0, 539, 540, 5, 97, 0, 0, 540, 541, 5,
		108, 0, 0, 541, 542, 5, 115, 0, 0, 542, 554, 5, 101, 0, 0, 543, 544, 5,
		70, 0, 0, 544, 545, 5, 97, 0, 0, 545, 546, 5, 108, 0, 0, 546, 547, 5, 115,
		0, 0, 547, 554, 5, 101, 0, 0, 548, 549, 5, 70, 0, 0, 549, 550, 5, 65, 0,
		0, 550, 551, 5, 76, 0, 0, 551, 552, 5, 83, 0, 0, 552, 554, 5, 69, 0, 0,
		553, 526, 1, 0, 0, 0, 553, 530, 1, 0, 0, 0, 553, 534, 1, 0, 0, 0, 553,
		538, 1, 0, 0, 0, 553, 543, 1, 0, 0, 0, 553, 548, 1, 0, 0, 0, 554, 82, 1,
		0, 0, 0, 555, 560, 3, 111, 55, 0, 556, 560, 3, 113, 56, 0, 557, 560, 3,
		115, 57, 0, 558, 560, 3, 109, 54, 0, 559, 555, 1, 0, 0, 0, 559, 556, 1,
		0, 0, 0, 559, 557, 1, 0, 0, 0, 559, 558, 1, 0, 0, 0, 560, 84, 1, 0, 0,
		0, 561, 564, 3, 127, 63, 0, 562, 564, 3, 129, 64, 0, 563, 561, 1, 0, 0,
		0, 563, 562, 1, 0, 0, 0, 564, 86, 1, 0, 0, 0, 565, 570, 3, 105, 52, 0,
		566, 569, 3, 105, 52, 0, 567, 569, 3, 107, 53, 0, 568, 566, 1, 0, 0, 0,
		568, 567, 1, 0, 0, 0, 569, 572, 1, 0, 0, 0, 570, 568, 1, 0, 0, 0, 570,
		571, 1, 0, 0, 0, 571, 88, 1, 0, 0, 0, 572, 570, 1, 0, 0, 0, 573, 574, 5,
		36, 0, 0, 574, 575, 5, 109, 0, 0, 575, 576, 5, 101, 0, 0, 576, 577, 5,
		116, 0, 0, 577, 578, 5, 97, 0, 0, 578, 90, 1, 0, 0, 0, 579, 581, 3, 95,
		47, 0, 580, 579, 1, 0, 0, 0, 580, 581, 1, 0, 0, 0, 581, 592, 1, 0, 0, 0,
		582, 584, 5, 34, 0, 0, 583, 585, 3, 97, 48, 0, 584, 583, 1, 0, 0, 0, 584,
		585, 1, 0, 0, 0, 585, 586, 1, 0, 0, 0, 586, 593, 5, 34, 0, 0, 587, 589,
		5, 39, 0, 0, 588, 590, 3, 99, 49, 0, 589, 588, 1, 0, 0, 0, 589, 590, 1,
		0, 0, 0, 590, 591, 1, 0, 0, 0, 591, 593, 5, 39, 0, 0, 592, 582, 1, 0, 0,
		0, 592, 587, 1, 0, 0, 0, 593, 92, 1, 0, 0, 0, 594, 597, 3, 87, 43, 0, 595,
		597, 3, 89, 44, 0, 596, 594, 1, 0, 0, 0, 596, 595, 1, 0, 0, 0, 597, 605,
		1, 0, 0, 0, 598, 601, 5, 91, 0, 0, 599, 602, 3, 91, 45, 0, 600, 602, 3,
		111, 55, 0, 601, 599, 1, 0, 0, 0, 601, 600, 1, 0, 0, 0, 602, 603, 1, 0,
		0, 0, 603, 604, 5, 93, 0, 0, 604, 606, 1, 0, 0, 0, 605, 598, 1, 0, 0, 0,
		606, 607, 1, 0, 0, 0, 607, 605, 1, 0, 0, 0, 607, 608, 1, 0, 0, 0, 608,
		94, 1, 0, 0, 0, 609, 610, 5, 117, 0, 0, 610, 613, 5, 56, 0, 0, 611, 613,
		7, 0, 0, 0, 612, 609, 1, 0, 0, 0, 612, 611, 1, 0, 0, 0, 613, 96, 1, 0,
		0, 0, 614, 616, 3, 101, 50, 0, 615, 614, 1, 0, 0, 0, 616, 617, 1, 0, 0,
		0, 617, 615, 1, 0, 0, 0, 617, 618, 1, 0, 0, 0, 618, 98, 1, 0, 0, 0, 619,
		621, 3, 103, 51, 0, 620, 619, 1, 0, 0, 0, 621, 622, 1, 0, 0, 0, 622, 620,
		1, 0, 0, 0, 622, 623, 1, 0, 0, 0, 623, 100, 1, 0, 0, 0, 624, 632, 8, 1,
		0, 0, 625, 632, 3, 143, 71, 0, 626, 627, 5, 92, 0, 0, 627, 632, 5, 10,
		0, 0, 628, 629, 5, 92, 0, 0, 629, 630, 5, 13, 0, 0, 630, 632, 5, 10, 0,
		0, 631, 624, 1, 0, 0, 0, 631, 625, 1, 0, 0, 0, 631, 626, 1, 0, 0, 0, 631,
		628, 1, 0, 0, 0, 632, 102, 1, 0, 0, 0, 633, 641, 8, 2, 0, 0, 634, 641,
		3, 143, 71, 0, 635, 636, 5, 92, 0, 0, 636, 641, 5, 10, 0, 0, 637, 638,
		5, 92, 0, 0, 638, 639, 5, 13, 0, 0, 639, 641, 5, 10, 0, 0, 640, 633, 1,
		0, 0, 0, 640, 634, 1, 0, 0, 0, 640, 635, 1, 0, 0, 0, 640, 637, 1, 0, 0,
		0, 641, 104, 1, 0, 0, 0, 642, 643, 7, 3, 0, 0, 643, 106, 1, 0, 0, 0, 644,
		645, 7, 4, 0, 0, 645, 108, 1, 0, 0, 0, 646, 647, 5, 48, 0, 0, 647, 649,
		7, 5, 0, 0, 648, 650, 7, 6, 0, 0, 649, 648, 1, 0, 0, 0, 650, 651, 1, 0,
		0, 0, 651, 649, 1, 0, 0, 0, 651, 652, 1, 0, 0, 0, 652, 110, 1, 0, 0, 0,
		653, 657, 3, 117, 58, 0, 654, 656, 3, 107, 53, 0, 655, 654, 1, 0, 0, 0,
		656, 659, 1, 0, 0, 0, 657, 655, 1, 0, 0, 0, 657, 658, 1, 0, 0, 0, 658,
		662, 1, 0, 0, 0, 659, 657, 1, 0, 0, 0, 660, 662, 5, 48, 0, 0, 661, 653,
		1, 0, 0, 0, 661, 660, 1, 0, 0, 0, 662, 112, 1, 0, 0, 0, 663, 667, 5, 48,
		0, 0, 664, 666, 3, 119, 59, 0, 665, 664, 1, 0, 0, 0, 666, 669, 1, 0, 0,
		0, 667, 665, 1, 0, 0, 0, 667, 668, 1, 0, 0, 0, 668, 114, 1, 0, 0, 0, 669,
		667, 1, 0, 0, 0, 670, 671, 5, 48, 0, 0, 671, 672, 7, 7, 0, 0, 672, 673,
		3, 139, 69, 0, 673, 116, 1, 0, 0, 0, 674, 675, 7, 8, 0, 0, 675, 118, 1,
		0, 0, 0, 676, 677, 7, 9, 0, 0, 677, 120, 1, 0, 0, 0, 678, 679, 7, 10, 0,
		0, 679, 122, 1, 0, 0, 0, 680, 681, 3, 121, 60, 0, 681, 682, 3, 121, 60,
		0, 682, 683, 3, 121, 60, 0, 683, 684, 3, 121, 60, 0, 684, 124, 1, 0, 0,
		0, 685, 686, 5, 92, 0, 0, 686, 687, 5, 117, 0, 0, 687, 688, 1, 0, 0, 0,
		688, 696, 3, 123, 61, 0, 689, 690, 5, 92, 0, 0, 690, 691, 5, 85, 0, 0,
		691, 692, 1, 0, 0, 0, 692, 693, 3, 123, 61, 0, 693, 694, 3, 123, 61, 0,
		694, 696, 1, 0, 0, 0, 695, 685, 1, 0, 0, 0, 695, 689, 1, 0, 0, 0, 696,
		126, 1, 0, 0, 0, 697, 699, 3, 131, 65, 0, 698, 700, 3, 133, 66, 0, 699,
		698, 1, 0, 0, 0, 699, 700, 1, 0, 0, 0, 700, 705, 1, 0, 0, 0, 701, 702,
		3, 135, 67, 0, 702, 703, 3, 133, 66, 0, 703, 705, 1, 0, 0, 0, 704, 697,
		1, 0, 0, 0, 704, 701, 1, 0, 0, 0, 705, 128, 1, 0, 0, 0, 706, 707, 5, 48,
		0, 0, 707, 710, 7, 7, 0, 0, 708, 711, 3, 137, 68, 0, 709, 711, 3, 139,
		69, 0, 710, 708, 1, 0, 0, 0, 710, 709, 1, 0, 0, 0, 711, 712, 1, 0, 0, 0,
		712, 713, 3, 141, 70, 0, 713, 130, 1, 0, 0, 0, 714, 716, 3, 135, 67, 0,
		715, 714, 1, 0, 0, 0, 715, 716, 1, 0, 0, 0, 716, 717, 1, 0, 0, 0, 717,
		718, 5, 46, 0, 0, 718, 723, 3, 135, 67, 0, 719, 720, 3, 135, 67, 0, 720,
		721, 5, 46, 0, 0, 721, 723, 1, 0, 0, 0, 722, 715, 1, 0, 0, 0, 722, 719,
		1, 0, 0, 0, 723, 132, 1, 0, 0, 0, 724, 726, 7, 11, 0, 0, 725, 727, 7, 12,
		0, 0, 726, 725, 1, 0, 0, 0, 726, 727, 1, 0, 0, 0, 727, 728, 1, 0, 0, 0,
		728, 729, 3, 135, 67, 0, 729, 134, 1, 0, 0, 0, 730, 732, 3, 107, 53, 0,
		731, 730, 1, 0, 0, 0, 732, 733, 1, 0, 0, 0, 733, 731, 1, 0, 0, 0, 733,
		734, 1, 0, 0, 0, 734, 136, 1, 0, 0, 0, 735, 737, 3, 139, 69, 0, 736, 735,
		1, 0, 0, 0, 736, 737, 1, 0, 0, 0, 737, 738, 1, 0, 0, 0, 738, 739, 5, 46,
		0, 0, 739, 744, 3, 139, 69, 0, 740, 741, 3, 139, 69, 0, 741, 742, 5, 46,
		0, 0, 742, 744, 1, 0, 0, 0, 743, 736, 1, 0, 0, 0, 743, 740, 1, 0, 0, 0,
		744, 138, 1, 0, 0, 0, 745, 747, 3, 121, 60, 0, 746, 745, 1, 0, 0, 0, 747,
		748, 1, 0, 0, 0, 748, 746, 1, 0, 0, 0, 748, 749, 1, 0, 0, 0, 749, 140,
		1, 0, 0, 0, 750, 752, 7, 13, 0, 0, 751, 753, 7, 12, 0, 0, 752, 751, 1,
		0, 0, 0, 752, 753, 1, 0, 0, 0, 753, 754, 1, 0, 0, 0, 754, 755, 3, 135,
		67, 0, 755, 142, 1, 0, 0, 0, 756, 757, 5, 92, 0, 0, 757, 772, 7, 14, 0,
		0, 758, 759, 5, 92, 0, 0, 759, 761, 3, 119, 59, 0, 760, 762, 3, 119, 59,
		0, 761, 760, 1, 0, 0, 0, 761, 762, 1, 0, 0, 0, 762, 764, 1, 0, 0, 0, 763,
		765, 3, 119, 59, 0, 764, 763, 1, 0, 0, 0, 764, 765, 1, 0, 0, 0, 765, 772,
		1, 0, 0, 0, 766, 767, 5, 92, 0, 0, 767, 768, 5, 120, 0, 0, 768, 769, 1,
		0, 0, 0, 769, 772, 3, 139, 69, 0, 770, 772, 3, 125, 62, 0, 771, 756, 1,
		0, 0, 0, 771, 758, 1, 0, 0, 0, 771, 766, 1, 0, 0, 0, 771, 770, 1, 0, 0,
		0, 772, 144, 1, 0, 0, 0, 773, 775, 7, 15, 0, 0, 774, 773, 1, 0, 0, 0, 775,
		776, 1, 0, 0, 0, 776, 774, 1, 0, 0, 0, 776, 777, 1, 0, 0, 0, 777, 778,
		1, 0, 0, 0, 778, 779, 6, 72, 0, 0, 779, 146, 1, 0, 0, 0, 780, 782, 5, 13,
		0, 0, 781, 783, 5, 10, 0, 0, 782, 781, 1, 0, 0, 0, 782, 783, 1, 0, 0, 0,
		783, 786, 1, 0, 0, 0, 784, 786, 5, 10, 0, 0, 785, 780, 1, 0, 0, 0, 785,
		784, 1, 0, 0, 0, 786, 787, 1, 0, 0, 0, 787, 788, 6, 73, 0, 0, 788, 148,
		1, 0, 0, 0, 56, 0, 187, 201, 230, 262, 268, 276, 282, 287, 289, 320, 356,
		392, 422, 460, 498, 524, 553, 559, 563, 568, 570, 580, 584, 589, 592, 596,
		601, 607, 612, 617, 622, 631, 640, 651, 657, 661, 667, 695, 699, 704, 710,
		715, 722, 726, 733, 736, 743, 748, 752, 761, 764, 771, 776, 782, 785, 1,
		6, 0, 0,
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

// PlanLexerInit initializes any static state used to implement PlanLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPlanLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PlanLexerInit() {
	staticData := &PlanLexerLexerStaticData
	staticData.once.Do(planlexerLexerInit)
}

// NewPlanLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPlanLexer(input antlr.CharStream) *PlanLexer {
	PlanLexerInit()
	l := new(PlanLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PlanLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Plan.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PlanLexer tokens.
const (
	PlanLexerT__0             = 1
	PlanLexerT__1             = 2
	PlanLexerT__2             = 3
	PlanLexerT__3             = 4
	PlanLexerT__4             = 5
	PlanLexerLBRACE           = 6
	PlanLexerRBRACE           = 7
	PlanLexerLT               = 8
	PlanLexerLE               = 9
	PlanLexerGT               = 10
	PlanLexerGE               = 11
	PlanLexerEQ               = 12
	PlanLexerNE               = 13
	PlanLexerLIKE             = 14
	PlanLexerEXISTS           = 15
	PlanLexerTEXTMATCH        = 16
	PlanLexerADD              = 17
	PlanLexerSUB              = 18
	PlanLexerMUL              = 19
	PlanLexerDIV              = 20
	PlanLexerMOD              = 21
	PlanLexerPOW              = 22
	PlanLexerSHL              = 23
	PlanLexerSHR              = 24
	PlanLexerBAND             = 25
	PlanLexerBOR              = 26
	PlanLexerBXOR             = 27
	PlanLexerAND              = 28
	PlanLexerOR               = 29
	PlanLexerBNOT             = 30
	PlanLexerNOT              = 31
	PlanLexerIN               = 32
	PlanLexerEmptyArray       = 33
	PlanLexerJSONContains     = 34
	PlanLexerJSONContainsAll  = 35
	PlanLexerJSONContainsAny  = 36
	PlanLexerArrayContains    = 37
	PlanLexerArrayContainsAll = 38
	PlanLexerArrayContainsAny = 39
	PlanLexerArrayLength      = 40
	PlanLexerBooleanConstant  = 41
	PlanLexerIntegerConstant  = 42
	PlanLexerFloatingConstant = 43
	PlanLexerIdentifier       = 44
	PlanLexerMeta             = 45
	PlanLexerStringLiteral    = 46
	PlanLexerJSONIdentifier   = 47
	PlanLexerWhitespace       = 48
	PlanLexerNewline          = 49
)
