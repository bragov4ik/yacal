package lexer_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
)

type testCase struct {
	input  string
	tokens []interface{}
}

var (
	LB   = tok.LBrace{}
	RB   = tok.RBrace{}
	NULL = tok.Null{}
)

func id(s string) tok.Ident { return tok.NewIdent(s) }

func readTokens(t *testing.T, l *lex.Lexer) ([]interface{}, error) {
	tokens := []interface{}{}

	for ty, at, token := l.Lex(); ty != tok.EOF; ty, at, token = l.Lex() {
		at += 1
		if err, isErr := token.(error); isErr {
			return nil, pp.Errorf("Error happened at %v: %v\n", at, err)
		}

		// in order to make explicit copy of empty interface
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func parsesTo(input string, tokens ...interface{}) testCase {
	return testCase{input, tokens}
}

func TestElements(t *testing.T) {
	tests := []testCase{
		parsesTo("1", 1),
		parsesTo("0123", 123),
		parsesTo("+1", 1),
		parsesTo("-1", -1),

		parsesTo("5.1", 5.1),
		parsesTo("+5.1", 5.1),
		parsesTo("-5.1", -5.1),

		parsesTo("true", true),
		parsesTo("false", false),
		parsesTo("null", NULL),

		parsesTo("'a'", 'a'),
		parsesTo("'\\\\'", '\\'),
		parsesTo("'\"'", '"'),
		parsesTo("'\\''", '\''),

		parsesTo("\"\"", ""),
		parsesTo("\"alola\"", "alola"),
		parsesTo("\"al\\\"ola\"", "al\"ola"),
		parsesTo("\"\"", ""),
		parsesTo("\"alola\"", "alola"),
		parsesTo("\"al\\\"ola\"", "al\"ola"),
		parsesTo("\"'''\"", "'''"),

		parsesTo("a4", id("a4")),
		parsesTo("()", LB, RB),
		parsesTo("(     )", LB, RB),
		parsesTo("(+ 1 2)", LB, id("+"), 1, 2, RB),
		parsesTo("(+12)", LB, 12, RB),
		parsesTo("(setq x 2)", LB, id("setq"), id("x"), 2, RB),
		parsesTo("(setq x (quote y))", LB, id("setq"), id("x"), LB, id("quote"), id("y"), RB, RB),
		parsesTo("(repeat \":kae:\" 1000)", LB, id("repeat"), ":kae:", 1000, RB),
		parsesTo("(append \"a\\\"bo\" \"ba\\\"\")", LB, id("append"), "a\"bo", "ba\"", RB),
		parsesTo("(setq x(+ 2 3))", LB, id("setq"), id("x"), LB, id("+"), 2, 3, RB, RB),
	}

	for _, tc := range tests {
		l := lexer.New(lex.NewFile("tmp", strings.NewReader(tc.input)))
		tokens, err := readTokens(t, l)
		if err != nil {
			t.Fatalf("Discover error (%v) while parsing %v", err, tc.input)
		}

		if !reflect.DeepEqual(tokens, tc.tokens) {
			t.Fatal(pp.Sprintf(
				"%v should be tokenized to\n%v\nbut got\n%v",
				tc.input, tc.tokens, tokens,
			))
		}
	}
}
