package main

import (
	"strings"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
)

func main() {
	l := lexer.New(lex.NewFile("tmp", strings.NewReader("(+ 1 2) (x 1 2 (1 2 3))")))
	for token := l.Eat(); token.Ty != tok.EOF; token = l.Eat() {
		if err, isErr := token.Value.(error); isErr {
			pp.Fatalf("Error happened at %v: %v\n", token.At, err)
		}
		pp.Printf("token of type %v at %v with value %v\n", token.Ty, token.At, token.Value)
	}
}
