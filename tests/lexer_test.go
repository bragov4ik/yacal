package lexer_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/db47h/lex"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
)

func ReadTokens(l *lex.Lexer) (tokens []lex.Token) {
	tokens = []lex.Token{}
	for ty, _, _ := l.Lex(); ty != tok.EOF; ty, _, _ = l.Lex() {
		tokens = append(tokens, ty)
	}
	return
}

type testCase struct {
	input  string
	tokens []lex.Token
}

func TestElements(t *testing.T) {
	for _, tc := range []testCase{
		{input: "1", tokens: []lex.Token{tok.INT}},
		{input: "0123", tokens: []lex.Token{tok.INT}},
		{input: "+1", tokens: []lex.Token{tok.INT}},
		{input: "-1", tokens: []lex.Token{tok.INT}},

		{input: "5.1", tokens: []lex.Token{tok.REAL}},
		{input: "+5.1", tokens: []lex.Token{tok.REAL}},
		{input: "-5.1", tokens: []lex.Token{tok.REAL}},

		{input: "true", tokens: []lex.Token{tok.BOOL}},
		{input: "false", tokens: []lex.Token{tok.BOOL}},

		{input: "null", tokens: []lex.Token{tok.NULL}},

		{input: "a4", tokens: []lex.Token{tok.IDENT}},

		{input: "\"\"", tokens: []lex.Token{tok.STRING}},
		{input: "\"alola\"", tokens: []lex.Token{tok.STRING}},
		{input: "\"al\\\"ola\"", tokens: []lex.Token{tok.STRING}},

		{input: "()", tokens: []lex.Token{tok.LBRACE, tok.RBRACE}},
		{input: "(     )", tokens: []lex.Token{tok.LBRACE, tok.RBRACE}},
		{input: "(+ 1 2)", tokens: []lex.Token{tok.LBRACE, tok.IDENT, tok.INT, tok.INT, tok.RBRACE}},
		{input: "(+12)", tokens: []lex.Token{tok.LBRACE, tok.INT, tok.RBRACE}},
		{input: "(setq x 2)", tokens: []lex.Token{tok.LBRACE, tok.IDENT, tok.IDENT, tok.INT, tok.RBRACE}},
		{input: "(repeat \":kae:\" 1000)", tokens: []lex.Token{tok.LBRACE, tok.IDENT, tok.STRING, tok.INT, tok.RBRACE}},
		{input: "(append \"a\\\"bo\" \"ba\\\"\")", tokens: []lex.Token{tok.LBRACE, tok.IDENT, tok.STRING, tok.STRING, tok.RBRACE}},
	} {
		l := lexer.New(lex.NewFile("tmp", strings.NewReader(tc.input)))
		tokens := ReadTokens(l)
		if !reflect.DeepEqual(tokens, tc.tokens) {
			t.Errorf("string %v should be tokenized to %v, but got %v", tc.input, tc.tokens, tokens)
		}
	}
}
