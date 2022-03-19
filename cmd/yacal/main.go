package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/db47h/lex"
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/interpreter"
	"github.com/bragov4ik/yacal/pkg/interpreter/builtin"
	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/parser"
)

func runRepl() {
	reader := bufio.NewReader(os.Stdin)
	i := interpreter.New()

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			pp.Fatalf("Failed to read from stdin: %v", err)
		}

		f := lex.NewFile("<stdin>", strings.NewReader(text))
		l := lexer.New(f)
		p := parser.New(l)
		ast, err := p.Parse()

		if err != nil {
			panic(fmt.Errorf("syntax error: %v", err))
		}

		ast = interpreter.TreeTraversal(ast)

		for _, st := range ast {
			v, err := i.Eval(st)
			if err != nil {
				panic(fmt.Sprintf("Failed to interpret %v: %v", v, err))
			}
			fmt.Println(builtin.ToString(v))
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		runRepl()
		return
	}

	i := interpreter.New()

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

		ast = interpreter.TreeTraversal(ast)

		for _, st := range ast {
			v, err := i.Eval(st)
			if err != nil {
				panic(fmt.Sprintf("Failed to interpret %v: %v", v, err))
			}
		}
	}
}
