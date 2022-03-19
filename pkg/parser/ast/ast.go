package ast

type (
	List []interface{}
	Cons struct {
		Val  interface{}
		Next interface{}
	}
	Empty struct{}
	Atom  struct{ Val string }
	Null  struct{}
)
