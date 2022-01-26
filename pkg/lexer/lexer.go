package lexer

type Span struct {
	lineNum          int64
	posBegin, posEnd int
}

type Token interface {
	Span() Span
}

func tokenize(source_text string) []Token {
	result := make([]Token, 0, 0)
	return result
}
