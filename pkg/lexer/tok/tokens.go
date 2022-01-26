package tok

import (
	"github.com/db47h/lex"
)

const (
	EOF lex.Token = iota

	LBRACE
	RBRACE
	// Native string type in golang
	IDENT
	// Native int type in golang
	INT
	// Native float64 type in golang
	REAL
	// Native bool type in golang
	BOOL
	// Native rune type in golang
	LETTER
	// String
	STRING
	NULL
)

type LBrace struct{}
type RBrace struct{}
type Quote struct{}
type Null struct{}
type Eof struct{}
