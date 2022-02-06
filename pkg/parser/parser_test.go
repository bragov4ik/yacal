package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/parser"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

type testCase struct {
	input string
	prog  []interface{}
}

func list(list ...interface{}) ast.List                    { return list }
func atom(s string) ast.Atom                               { return ast.Atom{Val: s} }
func parsesTo(input string, elems ...interface{}) testCase { return testCase{input, elems} }

func TestElements(t *testing.T) {
	tests := []testCase{
		parsesTo("(+ 1 2)", list(atom("+"), 1, 2)),
		parsesTo("(+12)", list(12)),
		parsesTo("(setq x 2)", list(atom("setq"), atom("x"), 2)),
		parsesTo("(setq x (quote y))", list(atom("setq"), atom("x"), list(atom("quote"), atom("y")))),
		parsesTo("(repeat \":kae:\" 1000)", list(atom("repeat"), ":kae:", 1000)),
		parsesTo("(setq x(+ 2 3))", list(atom("setq"), atom("x"), list(atom("+"), 2, 3))),
	}

	for i, tc := range tests {
		prog, err := parser.New(lexer.New(lex.NewFile("tmp", strings.NewReader(tc.input)))).Parse()
		if err != nil {
			t.Fatalf("Discover error (%v) while parsing test %v: %v", err, i, tc.input)
		}

		if !reflect.DeepEqual(prog, tc.prog) {
			t.Fatal(pp.Sprintf(
				"%v should be made into this ast\n%v\nbut got\n%v",
				tc.input, tc.prog, prog,
			))
		}
	}
}
