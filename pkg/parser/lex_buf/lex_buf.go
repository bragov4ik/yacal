package lex_buf

import "github.com/db47h/lex"

type LexResult struct {
	Ty    lex.Token
	At    int
	Token interface{}
}

type LexerBuf struct {
	lexer   *lex.Lexer
	curLex  LexResult
	nextLex LexResult
}

func New(lexer *lex.Lexer) *LexerBuf {
	l := &LexerBuf{lexer, LexResult{}, LexResult{}}
	// one step so that first Next() gives the first token
	l.Next()
	return l
}

func (l *LexerBuf) Next() LexResult {
	ty, at, token := l.lexer.Lex()
	at += 1
	l.curLex = l.nextLex
	l.nextLex = LexResult{ty, at, token}
	return l.curLex
}

func (l *LexerBuf) Current() LexResult {
	return l.curLex
}

func (l *LexerBuf) Peek() LexResult {
	return l.nextLex
}
