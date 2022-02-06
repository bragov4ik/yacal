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
	"github.com/bragov4ik/yacal/pkg/parser/lex_buf"
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
		parsesTo("1", &tree.Program{Elements: []tree.Node{&tree.Literal{Value: 1}}}),
		parsesTo("+5.1", &tree.Program{Elements: []tree.Node{&tree.Literal{Value: 5.1}}}),
		parsesTo("true", &tree.Program{Elements: []tree.Node{&tree.Literal{Value: true}}}),
		parsesTo("null", &tree.Program{Elements: []tree.Node{&tree.Literal{Value: NULL}}}),
		parsesTo("'a'", &tree.Program{Elements: []tree.Node{&tree.Literal{Value: 'a'}}}),
		parsesTo("\"alola\"", &tree.Program{Elements: []tree.Node{&tree.Literal{Value: "alola"}}}),

		parsesTo("a4", &tree.Program{Elements: []tree.Node{&tree.Atom{Ident: "a4"}}}),
		parsesTo("(+ 1 2)", &tree.Program{Elements: []tree.Node{&tree.List{Elements: []tree.Node{&tree.Atom{Ident: "+"}, &tree.Literal{Value: 1}, &tree.Literal{Value: 2}}}}}),
		parsesTo("(+12)", &tree.Program{Elements: []tree.Node{&tree.List{Elements: []tree.Node{&tree.Literal{Value: 12}}}}}),
		parsesTo("(setq x (quote y))", &tree.Program{Elements: []tree.Node{&tree.List{Elements: []tree.Node{&tree.Atom{Ident: "setq"}, &tree.Atom{Ident: "x"}, &tree.List{Elements: []tree.Node{&tree.Atom{Ident: "quote"}, &tree.Atom{Ident: "y"}}}}}}}),
		parsesTo("(setq x(+ 2 3))", &tree.Program{Elements: []tree.Node{&tree.List{Elements: []tree.Node{&tree.Atom{Ident: "setq"}, &tree.Atom{Ident: "x"}, &tree.List{Elements: []tree.Node{&tree.Atom{Ident: "+"}, &tree.Literal{Value: 2}, &tree.Literal{Value: 3}}}}}}}),
		parsesTo("(setq x(+ 2 3))", &tree.Program{Elements: []tree.Node{&tree.List{Elements: []tree.Node{&tree.Atom{Ident: "setq"}, &tree.Atom{Ident: "x"}, &tree.List{Elements: []tree.Node{&tree.Atom{Ident: "+"}, &tree.Literal{Value: 2}, &tree.Literal{Value: 3}}}}}}}),
		parsesTo(`(setq x 5)
		(setq y (plus 1 2))
		(setq z null)`, &tree.Program{Elements: []tree.Node{
			&tree.List{Elements: []tree.Node{&tree.Atom{Ident: "setq"}, &tree.Atom{Ident: "x"}, &tree.Literal{Value: 5}}},
			&tree.List{Elements: []tree.Node{&tree.Atom{Ident: "setq"}, &tree.Atom{Ident: "y"}, &tree.List{Elements: []tree.Node{&tree.Atom{Ident: "plus"}, &tree.Literal{Value: 1}, &tree.Literal{Value: 2}}}}},
			&tree.List{Elements: []tree.Node{&tree.Atom{Ident: "setq"}, &tree.Atom{Ident: "z"}, &tree.Literal{Value: NULL}}}}}),
	}

	for _, tc := range tests {
		l := lexer.New(lex.NewFile("tmp", strings.NewReader(tc.input)))
		lb := lex_buf.New(l)
		tree, err := parser.ParseProgram(lb)
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
