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
func NextToken(l *lex.Lexer) tokenInfo {
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

func ParseList(curTok tokenInfo, l *lex.Lexer) (*tree.List, error) {
	var tokenInfo tokenInfo
	// (
	if curTok.ty != tok.LBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", curTok.at, "List should start with '('")
	}

	// First element (mandatory)
	elements := make([]tree.Node, 0)
	tokenInfo = NextToken(l)
	if tokenInfo.ty == tok.RBRACE {
		return nil, pp.Errorf("Error happened at %v: %v\n", tokenInfo.at, "List should be non-empty")
	}
	elem, err := ParseElement(tokenInfo, l)
	if err == nil {
		elements = append(elements, elem)
	} else {
		return nil, err
	}

	// Other elements until closing brace
	for tokenInfo = NextToken(l); tokenInfo.ty != tok.EOF; tokenInfo = NextToken(l) {
		if tokenInfo.ty == tok.RBRACE {
			break
		}
		elem, err := ParseElement(tokenInfo, l)
		if err == nil {
			elements = append(elements, elem)
		} else {
			return nil, err
		}
	}
	return &tree.List{Elements: elements}, nil
}

func ParseElement(curTok tokenInfo, l *lex.Lexer) (tree.Node, error) {
	switch v := curTok.token.(type) {
	case tok.Ident:
		return parseAtom(v), nil
	case int, float64, bool, rune, string, tok.Null:
		return parseLiteral(v), nil
	case tok.LBrace:
		return ParseList(curTok, l)
	default:
		return nil, pp.Errorf("Error happened at %v: Unexpected token %v\n", curTok.at, curTok.token)
	}
}
