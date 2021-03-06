package lexer

import (
	"strconv"
	"unicode"

	"github.com/db47h/lex"

	"github.com/bragov4ik/yacal/pkg/lexer/tok"
)

type Lexer struct {
	lex *lex.Lexer
	cur Token
}

type Token struct {
	Ty    lex.Token
	At    int
	Value interface{}
}

func (l *Lexer) Peek() Token { return l.cur }

func (l *Lexer) Eat() Token {
	tok := l.cur
	ty, at, value := l.lex.Lex()
	// Weird bug in library. Everything starts at -1
	at += 1
	l.cur = Token{ty, at, value}
	return tok
}

func isAtomEnd(r rune) bool {
	return unicode.IsSpace(r) || r == lex.EOF || r == ')' || r == '('
}

func assertChar(l *lex.State, r rune) (rune, bool) {
	got := l.Next()
	return got, got != r
}

func skipComment(l *lex.State) lex.StateFn {
	for r := l.Next(); r != '\n'; r = l.Next() {
		if r == lex.EOF {
			l.Emit(l.Pos(), tok.EOF, tok.Eof{})
			return nil
		}
	}
	return nil
}

func readIdentInner(l *lex.State) string {
	value := ""
	for r := l.Peek(); !isAtomEnd(r); r = l.Peek() {
		value += string(r)
		l.Next()
	}
	return value
}

func readIdent(l *lex.State) lex.StateFn {
	l.StartToken(l.Pos())
	switch ident := readIdentInner(l); ident {
	case "true":
		l.Emit(l.TokenPos(), tok.BOOL, true)
	case "false":
		l.Emit(l.TokenPos(), tok.BOOL, false)
	case "null":
		l.Emit(l.TokenPos(), tok.NULL, tok.Null{})
	default:
		l.Emit(l.TokenPos(), tok.IDENT, tok.Ident{Val: ident})
	}
	return nil
}

func readInt(l *lex.State) string {
	value := ""
	for r := l.Peek(); unicode.IsDigit(r); r = l.Peek() {
		value += string(r)
		l.Next()
	}
	return value
}

func readNumber(l *lex.State) lex.StateFn {
	l.StartToken(l.Pos())
	// By default we are here if either first char is
	// [0-9] -- which means it is positive
	// +     -- which means it is positive
	// -     -- which means it is negative
	//
	// So there is a single case when it is positive
	mul := int64(1)
	switch r := l.Peek(); r {
	case '+':
		l.Next()
	case '-':
		l.Next()
		mul = -1
	}

	value := readInt(l)

	switch r := l.Peek(); {
	case r == '.':
		l.Next()
		value += string(r) + readInt(l)

		if r = l.Peek(); !isAtomEnd(r) {
			l.Errorf(l.TokenPos(), "Real numbers shouldn't end with `%v'", string(r))
			return nil
		}

		f, err := strconv.ParseFloat(value, 64)
		f *= float64(mul)
		if err != nil {
			l.Errorf(l.TokenPos(), "Error parsing real number: %v", err)
		} else {
			l.Emit(l.TokenPos(), tok.REAL, f)
		}
	case isAtomEnd(r):
		i, err := strconv.ParseInt(value, 10, 64)
		i *= mul
		new_i := int(i)
		if err != nil {
			l.Errorf(l.TokenPos(), "Error parsing integer: %v", err)
		} else {
			l.Emit(l.TokenPos(), tok.INT, new_i)
		}
	default:
		l.Errorf(l.TokenPos(), "Integers shouldn't end with `%v'", string(r))
	}

	return nil
}

func readString(l *lex.State) lex.StateFn {
	l.StartToken(l.Pos())
	if got, assert := assertChar(l, '"'); assert {
		l.Errorf(l.TokenPos(), "String should start with double quote, not with `%v'", string(got))
	}
	value := ""
	for r := l.Peek(); r != '"' && r != lex.EOF; r = l.Peek() {
		switch r {
		case '\\':
			l.Next()
			switch r = l.Peek(); r {
			case 'r':
				r = '\r'
			case 'n':
				r = '\n'
			case '\\':
				r = '\\'
			case '"':
				r = '"'
			case lex.EOF:
				l.Errorf(l.Pos(), "Expected character after backslash, but got EOF")
				return nil
			default:
				l.Errorf(l.Pos(), "Unknown escape char `%v'", string(r))
				return nil
			}
		}

		value += string(r)
		l.Next()
	}

	if l.Next() == lex.EOF {
		l.Errorf(l.Pos(), "Unexpected EOF. You should end string with double quote")
		return nil
	}
	l.Emit(l.TokenPos(), tok.STRING, value)
	return nil
}

func readTok(l *lex.State) lex.StateFn {
	l.StartToken(l.Pos())
	switch r := l.Next(); {
	case r == lex.EOF:
		l.Emit(l.TokenPos(), tok.EOF, tok.Eof{})
	case r == '/' && l.Peek() == '/':
		skipComment(l)
	case unicode.IsSpace(r):
		// skipping spaces
		break

	case r == '(':
		l.Emit(l.TokenPos(), tok.LBRACE, tok.LBrace{})
	case r == ')':
		l.Emit(l.TokenPos(), tok.RBRACE, tok.RBrace{})
	case r == '\'':
		l.Emit(l.TokenPos(), tok.QUOTE, tok.Quote{})
	case r == '"':
		if err := l.UnreadRune(); err != nil {
			l.Errorf(l.Pos(), "Unexpected error: %v", err)
		}
		return readString
	case unicode.IsDigit(r) || ((r == '+' || r == '-') && unicode.IsDigit(l.Peek())):
		if err := l.UnreadRune(); err != nil {
			l.Errorf(l.Pos(), "Unexpected error: %v", err)
		}
		return readNumber
	default:
		if err := l.UnreadRune(); err != nil {
			l.Errorf(l.Pos(), "Unexpected error: %v", err)
		}
		return readIdent
	}
	return nil
}

func New(f *lex.File) *Lexer {
	lex := lex.NewLexer(f, readTok)
	l := Lexer{lex, Token{}}
	l.Eat()
	return &l
}
