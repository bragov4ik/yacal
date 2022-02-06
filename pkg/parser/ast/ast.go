package ast

type (
	List   []interface{}
	Atom   struct{ Val string }
	String string
	Char   rune
	Int    int
	Real   float64
	Bool   bool
	Null   struct{}
)
