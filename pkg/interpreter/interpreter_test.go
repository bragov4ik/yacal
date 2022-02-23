package interpreter_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/interpreter"
	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/parser"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

type testCase struct {
	input  string
	output []interface{}
}

func list(list ...interface{}) ast.List                  { return list }
func atom(s string) ast.Atom                             { return ast.Atom{Val: s} }
func evalTo(input string, elems ...interface{}) testCase { return testCase{input, elems} }

func TestElements(t *testing.T) {
	tests := []testCase{
		evalTo("(+ 1 2)", 3),
		evalTo("(set x 2) x", nil, 2),
		evalTo("(set mul2 (lambda (x) (+ x x))) (mul2 10)", nil, 20),
		// evalTo("(car '(1 2 3))", list(atom("car"), list(atom("quote"), list(1, 2, 3)))),
		// evalTo(`(setq x 5) (setq y (plus 1 2)) (setq z null)`,
		// 	list(atom("setq"), atom("x"), 5),
		// 	list(atom("setq"), atom("y"), list(atom("plus"), 1, 2)),
		// 	list(atom("setq"), atom("z"), ast.Null{}),
		// ),
	}

	for i, tc := range tests {
		in := interpreter.New()
		prog, err := parser.New(lexer.New(lex.NewFile("tmp", strings.NewReader(tc.input)))).Parse()
		if err != nil {
			t.Fatalf("Discover error (%v) while parsing test %v: %v", err, i, tc.input)
		}

		out := []interface{}{}
		for _, st := range prog {
			if o, err := in.Eval(st); err != nil {
				t.Fatalf("Discover error (%v) while evaluating test %v: %v", err, i, tc.input)
			} else {
				out = append(out, o)
			}
		}

		if !reflect.DeepEqual(out, tc.output) {
			t.Fatal(pp.Sprintf(
				"%v should be evaluated into this\n%v\nbut got\n%v",
				tc.input, tc.output, out,
			))
		}
	}
}
