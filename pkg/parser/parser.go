package parser

import (
	"github.com/bragov4ik/yacal/pkg/lexer/buffer"
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
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

func parseList(l *buffer.LexerBuf) (*tree.List, error) {
	// Don't need to check left bracket, it is already checked by parseElement

	// Check first element
	if l.Peek().Ty == tok.RBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", l.Peek().At, "Unexpected ')', list should not be empty")
	} else if l.Peek().Ty == tok.EOF {
		return nil, pp.Errorf("Error happened at %v: %v\n", l.Peek().At, "Unexpected EOF, list is not closed")
	}

	elements := make([]tree.Node, 0)
	// First element (mandatory)
	elem, err := parseElement(l)
	if err == nil {
		elements = append(elements, elem)
	} else {
		return nil, err
	}

	// Other elements until EOF or closing brace
	for {
		if l.Peek().Ty == tok.RBRACE {
			l.Eat()
			break
		} else if l.Peek().Ty == tok.EOF {
			return nil, pp.Errorf("Error happened at %v: %v\n", l.Peek().At, "Unexpected EOF, list is not closed")
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

func parseElement(l *buffer.LexerBuf) (tree.Node, error) {
	nextElem := l.Eat()
	switch v := nextElem.Token.(type) {
	case tok.Ident:
		return parseAtom(v), nil
	case int, float64, bool, rune, string, tok.Null:
		return parseLiteral(v), nil
	case tok.LBrace:
		return parseList(l)
	default:
		return nil, pp.Errorf("Error happened at %v: Unexpected token %v\n", nextElem.At, nextElem.Token)
	}
}

func ParseProgram(l *buffer.LexerBuf) (*tree.Program, error) {
	// Check first element
	if l.Peek().Ty == tok.RBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", l.Peek().At, "Unexpected ')'")
	} else if l.Peek().Ty == tok.EOF {
		return nil, pp.Errorf("Error happened at %v: %v\n", l.Peek().At, "Unexpected EOF")
	}

	elements := make([]tree.Node, 0)
	// Add first element (mandatory)
	elem, err := parseElement(l)
	if err == nil {
		elements = append(elements, elem)
	} else {
		return nil, err
	}

	// Other elements until EOF
	for {
		if l.Peek().Ty == tok.EOF {
			break
		}
		elem, err := parseElement(l)
		if err == nil {
			elements = append(elements, elem)
		} else {
			return nil, err
		}
	}
	return &tree.Program{Elements: elements}, nil
}
