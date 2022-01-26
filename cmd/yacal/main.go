package main

import (
	"fmt"
	"strings"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
)

func main() {
	l := lexer.New(lex.NewFile("tmp", strings.NewReader("(+ 1 2) (x 1 2 (1 2 3))")))
	for ty, at, token := l.Lex(); ty != tok.EOF; ty, at, token = l.Lex() {
		// Weird bug in library. Everything starts at -1
		at += 1
		if err, isErr := token.(error); isErr {
			panic(fmt.Sprintf("Error happened at %v: %v\n", at, err))
		}
		pp.Printf("token of type %v at %v with value %v\n", ty, at, token)
	}
}
