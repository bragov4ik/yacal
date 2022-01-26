package lexer_test

import (
	"errors"
	"fmt"
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

func copyTok(tok interface{}) interface{} {
	v := reflect.Value{}
	reflect.Copy(v, reflect.ValueOf(tok))
	return v.Interface()
}

func readTokens(t *testing.T, l *lex.Lexer) ([]interface{}, error) {
	tokens := []interface{}{}
	for ty, at, token := l.Lex(); ty != tok.EOF; ty, _, _ = l.Lex() {
		at += 1
		if err, isErr := token.(error); isErr {
			return nil, errors.New(fmt.Sprintf("Error happened at %v: %v\n", at, err))
		}

		// in order to make explicit copy of empty interface
		tokens = append(tokens, copyTok(token))
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

		parsesTo("5.1", float64(5.1)),
		parsesTo("+5.1", float64(5.1)),
		parsesTo("-5.1", float64(-5.1)),

		parsesTo("true", true),
		parsesTo("false", false),
		parsesTo("null", tok.Null{}),

		parsesTo("a4", tok.NewIdent("a4")),
		parsesTo("()", tok.LBrace{}, tok.RBrace{}),
		parsesTo("(     )", tok.LBrace{}, tok.RBrace{}),
		parsesTo("(+ 1 2)", tok.LBrace{}, tok.NewIdent("+"), 1, 2, tok.RBrace{}),
		parsesTo("(+12)", tok.LBrace{}, 12, tok.RBrace{}),
		parsesTo("(setq x 2)", tok.LBrace{}, tok.NewIdent("setq"), tok.NewIdent("x"), 2, tok.RBrace{}),
	}

	for _, tc := range tests {
		l := lexer.New(lex.NewFile("tmp", strings.NewReader(tc.input)))
		tokens, err := readTokens(t, l)
		if err != nil {
			t.Fatalf("Discover error (%v) while parsing \"%v\"", err, tc.input)
		}
		pp.Printf("%v %v\n", tokens, tc.tokens)
		if !reflect.DeepEqual(tokens, tc.tokens) {
			t.Fatalf("\"%v\" should be tokenized to %v, but got %v", tc.input, tc.tokens, tokens)
		}
	}
}
