package parser

import (
	"github.com/bragov4ik/yacal/pkg/lexer/tok"
	"github.com/bragov4ik/yacal/pkg/parser/tree"
	"github.com/db47h/lex"
	"github.com/k0kubun/pp"
)

type tokenInfo struct {
	ty    lex.Token
	at    int
	token interface{}
}

// Obtain next token and report error if any
func nextToken(l *lex.Lexer) tokenInfo {
	ty, at, token := l.Lex()
	at += 1
	if err, isErr := token.(error); isErr {
		pp.Errorf("Error happened at %v: %v\n", at, err)
		// Somehow report error
		// return nil, pp.Errorf("Error happened at %v: %v\n", at, err)
	}
	return tokenInfo{ty, at, token}
}

// Terminals

func parseAtom(ident tok.Ident) *tree.Atom {
	return &tree.Atom{Ident: ident.String()}
}

func parseLiteral(value interface{}) *tree.Literal {
	return &tree.Literal{Value: value}
}

// Non-terminals

func parseList(curTok tokenInfo, l *lex.Lexer) (*tree.List, error) {
	var tokenInfo tokenInfo
	// (
	if curTok.ty != tok.LBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", curTok.at, "List should start with '('")
	}

	// First element (mandatory)
	elements := make([]interface{}, 0)
	tokenInfo = nextToken(l)
	if tokenInfo.ty == tok.RBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", tokenInfo.at, "List should be non-empty")
	}
	elem, err := parseElement(tokenInfo, l)
	if err == nil {
		elements = append(elements, elem)
	} else {
		return nil, err
	}

	// Other elements until closing brace
	for tokenInfo = nextToken(l); tokenInfo.ty != tok.EOF; tokenInfo = nextToken(l) {
		if tokenInfo.ty == tok.RBRACE {
			break
		}
		elem, err := parseElement(tokenInfo, l)
		if err == nil {
			elements = append(elements, elem)
		} else {
			return nil, err
		}
	}
	return &tree.List{Elements: elements}, nil
}

func parseElement(curTok tokenInfo, l *lex.Lexer) (interface{}, error) {
	switch v := curTok.token.(type) {
	case tok.Ident:
		return parseAtom(v), nil
	case int, float64, bool, rune, string, tok.Null:
		return parseLiteral(v), nil
	case tok.LBrace:
		return parseList(curTok, l)
	default:
		return nil, pp.Errorf("Error happened at %v: Unexpected token %v\n", curTok.at, curTok.token)
	}
}
