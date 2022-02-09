package parser

import (
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/lexer"
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

type Parser struct{ lex *lexer.Lexer }

func New(lex *lexer.Lexer) *Parser { return &Parser{lex} }

func (p *Parser) parseQuote() (ast.List, error) {
	if p.lex.Peek().Ty != tok.QUOTE {
		pp.Fatalf("Unexpected token %v", p.lex.Eat().Value)
	}
	p.lex.Eat()

	e, err := p.parseElem()
	if err != nil {
		return nil, pp.Errorf("Unexpected error decoding quote: %v", err)
	}
	return ast.List{ast.Atom{Val: "quote"}, e}, nil
}

func (p *Parser) parseList() (ast.List, error) {
	if p.lex.Peek().Ty != tok.LBRACE {
		pp.Fatalf("Unexpected token %v", p.lex.Eat().Value)
	}
	p.lex.Eat()

	l := ast.List{}

	for {
		if p.lex.Peek().Ty == tok.RBRACE {
			p.lex.Eat()
			return l, nil
		}
		if p.lex.Peek().Ty == tok.EOF {
			return l, pp.Errorf("Unexpected EOF while decoding list")
		}
		e, err := p.parseElem()
		if err != nil {
			return l, pp.Errorf("Unexpected error while decoding list: %v", err)
		}
		if e == nil {
			return l, pp.Errorf("Unexpected EOF while decoding list")
		}

		l = append(l, e)
	}
}

func (p *Parser) parseElem() (interface{}, error) {
	switch p.lex.Peek().Ty {
	case tok.LBRACE:
		return p.parseList()
	case tok.QUOTE:
		return p.parseQuote()
	case tok.IDENT:
		v := p.lex.Eat().Value.(tok.Ident).Val
		return ast.Atom{Val: v}, nil
	case tok.NULL:
		p.lex.Eat()
		return ast.Null{}, nil

	case tok.BOOL:
		fallthrough
	case tok.REAL:
		fallthrough
	case tok.STRING:
		fallthrough
	case tok.INT:
		return p.lex.Eat().Value, nil

	case tok.EOF:
		return nil, nil
	default:
		return nil, pp.Errorf("Expected some element, but got %v", p.lex.Eat().Value)
	}
}

func (p *Parser) Parse() ([]interface{}, error) {
	prog := []interface{}{}
	for {
		e, err := p.parseElem()
		if err != nil {
			return prog, err
		}
		if e == nil {
			return prog, nil
		}
		prog = append(prog, e)
	}
}
