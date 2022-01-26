package lexer

import (
	"strconv"
	"unicode"

	"github.com/db47h/lex"

	"github.com/bragov4ik/yacal/pkg/lexer/tok"
)

func skipComment(l *lex.State) lex.StateFn {
	for r := l.Next(); r != '\n'; r = l.Next() {
		if r == lex.EOF {
			l.Emit(l.Pos(), tok.EOF, &tok.Eof{})
			return nil
		}
	}
	return nil
}

func readIdentInner(l *lex.State) string {
	value := ""
	for r := l.Peek(); !unicode.IsSpace(r) && r != lex.EOF && r != ')'; r = l.Peek() {
		value += string(r)
		l.Next()
	}
	return value
}

func readIdent(l *lex.State) lex.StateFn {
	l.StartToken(l.Pos())
	switch ident := readIdentInner(l); ident {
	case "true":
		b := true
		l.Emit(l.TokenPos(), tok.BOOL, &b)
	case "false":
		b := false
		l.Emit(l.TokenPos(), tok.BOOL, &b)
	case "null":
		l.Emit(l.TokenPos(), tok.NULL, &tok.Null{})
	default:
		l.Emit(l.TokenPos(), tok.IDENT, &ident)
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
	value := readInt(l)

	switch r := l.Peek(); {
	case r == '.':
		l.Next()
		value += string(r) + readInt(l)

		if r = l.Peek(); !(r == lex.EOF || unicode.IsSpace(r) || r == ')') {
			l.Errorf(l.TokenPos(), "Real numbers should end with space, eof, or right bracket, not with '%v'", string(r))
			return nil
		}

		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			l.Errorf(l.TokenPos(), "Error parsing real number: %v", err)
		} else {
			l.Emit(l.TokenPos(), tok.REAL, &f)
		}
	case unicode.IsSpace(r) || r == lex.EOF || r == ')':
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			l.Errorf(l.TokenPos(), "Error parsing integer: %v", err)
		} else {
			l.Emit(l.TokenPos(), tok.INT, &i)
		}
	default:
		l.Errorf(l.TokenPos(), "Integers should end with space, eof, or right bracket, not with `%v'", string(r))
	}

	return nil
}

func readString(l *lex.State) lex.StateFn {
	l.StartToken(l.Pos())
	value := ""
	r := l.Next()
	if r != '"' {
		l.Errorf(l.TokenPos(), "String should start with \", not with '%v'", string(r))
	}
	for r := l.Next(); r != '"'; r = l.Next() {
		if r == lex.EOF {
			l.Errorf(l.TokenPos(), "Couldn't find end of string, reached EOF", string(r))
		}
		value += string(r)
	}
	l.Emit(l.TokenPos(), tok.STRING, &value)
	return nil
}

func readTok(l *lex.State) lex.StateFn {
	l.StartToken(l.Pos())
	switch r := l.Next(); {
	case r == lex.EOF:
		l.Emit(l.TokenPos(), tok.EOF, &tok.Eof{})
	case r == '(':
		l.Emit(l.TokenPos(), tok.LBRACE, &tok.LBrace{})
	case r == ')':
		l.Emit(l.TokenPos(), tok.RBRACE, &tok.RBrace{})
	case r == '\'':
		c := l.Next()
		if quote := l.Next(); quote != '\'' {
			l.Errorf(l.TokenPos(), "After character symbol should end with quote, not with `%v'", string(quote))
		}
		l.Emit(l.TokenPos(), tok.LETTER, &c)
	case r == '"':
		l.Backup()
		return readString
	// TODO: what if number starts with + or - sign?
	case unicode.IsDigit(r):
		l.Backup()
		return readNumber
	case r == '/' && l.Peek() == '/':
		skipComment(l)
	case unicode.IsSpace(r):
		// skipping spaces
		break
	default:
		l.Backup()
		return readIdent
	}
	return nil
}

func New(f *lex.File) *lex.Lexer { return lex.NewLexer(f, readTok) }
