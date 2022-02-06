package lex_buf

import "github.com/db47h/lex"

type LexResult struct {
	Ty    lex.Token
	At    int
	Token interface{}
}

type LexerBuf struct {
	lexer  *lex.Lexer
	curLex LexResult
}

func New(lexer *lex.Lexer) *LexerBuf {
	return &LexerBuf{lexer, LexResult{}}
}

func (l *LexerBuf) Next() LexResult {
	ty, at, token := l.lexer.Lex()
	at += 1
	l.curLex = LexResult{ty, at, token}
	return l.curLex
}

func (l *LexerBuf) Current() LexResult {
	return l.curLex
}
