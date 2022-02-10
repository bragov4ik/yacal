package main

import (
	"fmt"
	"os"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/parser"
)

func main() {
	global_ast := []interface{}{}

	for _, path := range os.Args[1:] {
		file, err := os.Open(path)
		if err != nil {
			panic(fmt.Errorf("failed to open file %v: %v", path, err))
		}

		f := lex.NewFile(path, file)
		l := lexer.New(f)
		p := parser.New(l)
		ast, err := p.Parse()

		if err != nil {
			panic(fmt.Errorf("got an error while building an ast for file %v: %v", path, err))
		}

		global_ast = append(global_ast, ast...)
	}

	pp.Printf("%v\n", global_ast)
}
