package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
	"github.com/bragov4ik/yacal/pkg/parser"
	"github.com/bragov4ik/yacal/pkg/parser/tree"
)

type testCase struct {
	input string
	tree  interface{}
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

func parsesTo(input string, tree interface{}) testCase {
	return testCase{input, tree}
}

func TestElements(t *testing.T) {
	tests := []testCase{
		// parsesTo("1", tree.Literal{1}),
		// parsesTo("0123", tree.Literal{123}),
		// parsesTo("+1", tree.Literal{1}),
		// parsesTo("-1", tree.Literal{-1}),

		// parsesTo("5.1", tree.Literal{5.1}),
		// parsesTo("+5.1", tree.Literal{5.1}),
		// parsesTo("-5.1", tree.Literal{-5.1}),

		// parsesTo("true", tree.Literal{true}),
		// parsesTo("false", tree.Literal{false}),
		// parsesTo("null", tree.Literal{NULL}),

		// parsesTo("'a'", tree.Literal{'a'}),
		// parsesTo("'\\\\'", tree.Literal{'\\'}),
		// parsesTo("'\"'", tree.Literal{'"'}),
		// parsesTo("'\\''", tree.Literal{'\''}),

		// parsesTo("\"\"", tree.Literal{""}),
		// parsesTo("\"alola\"", tree.Literal{"alola"}),
		// parsesTo("\"al\\\"ola\"", tree.Literal{"al\"ola"}),

		// parsesTo("a4", tree.Atom{"a4"}),
		parsesTo("(+ 1 2)", &tree.List{[]tree.Node{&tree.Atom{"+"}, &tree.Literal{1}, &tree.Literal{2}}}),
		parsesTo("(+12)", &tree.List{[]tree.Node{&tree.Literal{12}}}),
		parsesTo("(setq x 2)", &tree.List{[]tree.Node{&tree.Atom{"setq"}, &tree.Atom{"x"}, &tree.Literal{2}}}),
		parsesTo("(setq x (quote y))", &tree.List{[]tree.Node{&tree.Atom{"setq"}, &tree.Atom{"x"}, &tree.List{[]tree.Node{&tree.Atom{"quote"}, &tree.Atom{"y"}}}}}),
		parsesTo("(repeat \":kae:\" 1000)", &tree.List{[]tree.Node{&tree.Atom{"repeat"}, &tree.Literal{":kae:"}, &tree.Literal{1000}}}),
		parsesTo("(append \"a\\\"bo\" \"ba\\\"\")", &tree.List{[]tree.Node{&tree.Atom{"append"}, &tree.Literal{"a\"bo"}, &tree.Literal{"ba\""}}}),
		parsesTo("(setq x(+ 2 3))", &tree.List{[]tree.Node{&tree.Atom{"setq"}, &tree.Atom{"x"}, &tree.List{[]tree.Node{&tree.Atom{"+"}, &tree.Literal{2}, &tree.Literal{3}}}}}),
	}

	for _, tc := range tests {
		l := lexer.New(lex.NewFile("tmp", strings.NewReader(tc.input)))
		tree, err := parser.ParseList(parser.NextToken(l), l)
		if err != nil {
			t.Fatalf("Discover error (%v) while parsing %v", err, tc.input)
		}
		if !reflect.DeepEqual(tree, tc.tree) {
			t.Error(pp.Sprintf(
				"%v should be tokenized to\n%v\nbut got\n%v",
				tc.input, tc.tree, tree,
			))
		}
	}
}
