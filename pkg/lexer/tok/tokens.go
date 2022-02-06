package tok

import (
	"github.com/db47h/lex"
)

const (
	EOF lex.Token = iota

	LBRACE
	RBRACE
	IDENT
	NULL

	// Native int type in golang
	INT
	// Native float64 type in golang
	REAL
	// Native bool type in golang
	BOOL
	// Native rune type in golang
	LETTER
	// Native string type in golang
	STRING
)

type (
	LBrace struct{}
	RBrace struct{}
	Null   struct{}
	Eof    struct{}
	Ident  struct{ Val string }
)
