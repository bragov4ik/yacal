package interpreter_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/db47h/lex"

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
		evalTo("(/ 3 2 1)", 1),
		evalTo("(/ 3.0 2.0 2.0)", .75),
		evalTo("(* 3 2 1)", 6),
		evalTo("(set x 2) x", nil, 2),
		evalTo("(set mul2 (lambda (x) (+ x x))) (mul2 10)", nil, 20),
		evalTo("'(1 2 3)", list(1, 2, 3)),
		evalTo("(= (+ 2 2) 4)", true),
		evalTo("(set l (cons 4 '(3 2 1))) l (head l)", nil, list(4, 3, 2, 1), 4),
		evalTo("(tail '(4 3 2 1))", list(3, 2, 1)),
		evalTo("(and (and (isreal 1.0) (isbool true)) (and (isnull null) (islist '(1 2 3))))", true),
		evalTo("(xor (and (isreal 1.0) (isbool true)) (!= 4 3))", false),
		evalTo("(eval '(+ 1 2)) (eval 1)", 3, 1),
		evalTo(`(set max 
				(lambda (x y) (
					cond (> x y) x y)))
				(max 1 3)`, nil, 3),
		evalTo("(set a 1) (while (< a 10) (set a (* a 4))) a", nil, ast.Null{}, 16),
		evalTo("(isnull (while false (/ 1 0)))", true),
		evalTo(`(set lst '(1 2 3))
				(set x 404)
				(set y 404)
				(prog (x y) ((set x 333)
					(set y lst) 
				    (set x (+ x (head y)))
					x))
				x
				y`, nil, nil, nil, 334, 404, 404),
		evalTo("(= (toint \"123\") (toreal \"123.0000\"))", true),
		// evalTo("(car '(1 2 3))", list(atom("car"), list(atom("quote"), list(1, 2, 3)))),
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
			t.Fatal(fmt.Sprintf(
				"%v should be evaluated into this\n%v\nbut got\n%v",
				tc.input, tc.output, out,
			))
		}
	}
}
