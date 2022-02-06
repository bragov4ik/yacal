package buffer

import "github.com/db47h/lex"

type Token struct {
	Ty    lex.Token
	At    int
	Token interface{}
}

type LexerBuf struct {
	lexer   *lex.Lexer
	current Token
	next    Token
}

func New(lexer *lex.Lexer) *LexerBuf {
	l := &LexerBuf{lexer, Token{}, Token{}}
	// one step so that first Next() gives the first token
	l.Eat()
	return l
}

func (l *LexerBuf) Eat() Token {
	ty, at, token := l.lexer.Lex()
	at += 1
	l.current = l.next
	l.next = Token{ty, at, token}
	return l.current
}

func (l *LexerBuf) Peek() Token {
	return l.next
}
