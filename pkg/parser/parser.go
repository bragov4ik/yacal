package parser

import (
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
	"github.com/bragov4ik/yacal/pkg/parser/lex_buf"
	"github.com/bragov4ik/yacal/pkg/parser/tree"
	"github.com/k0kubun/pp"
)

// Terminals

func parseAtom(ident tok.Ident) *tree.Atom {
	return &tree.Atom{Ident: ident.String()}
}

func parseLiteral(value interface{}) *tree.Literal {
	return &tree.Literal{Value: value}
}

// Non-terminals

func parseList(l *lex_buf.LexerBuf) (*tree.List, error) {
	var tokenInfo lex_buf.LexResult
	// (
	if l.Current().Ty != tok.LBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", l.Current().At, "List should start with '('")
	}

	// First element (mandatory)
	elements := make([]tree.Node, 0)
	tokenInfo = l.Next()
	if tokenInfo.Ty == tok.RBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", tokenInfo.At, "List should be non-empty")
	}
	elem, err := parseElement(l)
	if err == nil {
		elements = append(elements, elem)
	} else {
		return nil, err
	}

	// Other elements until closing brace
	for tokenInfo = l.Next(); tokenInfo.Ty != tok.EOF; tokenInfo = l.Next() {
		if tokenInfo.Ty == tok.RBRACE {
			break
		}
		elem, err := parseElement(l)
		if err == nil {
			elements = append(elements, elem)
		} else {
			return nil, err
		}
	}
	return &tree.List{Elements: elements}, nil
}

func parseElement(l *lex_buf.LexerBuf) (tree.Node, error) {
	switch v := l.Current().Token.(type) {
	case tok.Ident:
		return parseAtom(v), nil
	case int, float64, bool, rune, string, tok.Null:
		return parseLiteral(v), nil
	case tok.LBrace:
		return parseList(l)
	default:
		return nil, pp.Errorf("Error happened at %v: Unexpected token %v\n", l.Current().At, l.Current().Token)
	}
}

func ParseProgram(l *lex_buf.LexerBuf) (*tree.Program, error) {
	// First element (mandatory)
	elements := make([]tree.Node, 0)
	tokenInfo := l.Next()
	if tokenInfo.Ty == tok.RBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", tokenInfo.At, "Program should be non-empty")
	}
	elem, err := parseElement(l)
	if err == nil {
		elements = append(elements, elem)
	} else {
		return nil, err
	}

	// Other elements until EOF
	for tokenInfo = l.Next(); tokenInfo.Ty != tok.EOF; tokenInfo = l.Next() {
		elem, err := parseElement(l)
		if err == nil {
			elements = append(elements, elem)
		} else {
			return nil, err
		}
	}
	return &tree.Program{Elements: elements}, nil
}
