package tree

type Node interface {
	node()
}

type Atom struct {
	Ident string
}

type Literal struct {
	Value interface{}
}

type List struct {
	Elements []Node
}

// Note: does not implement Node
type Program struct {
	Elements []Node
}

func (*Atom) node()    {}
func (*Literal) node() {}
func (*List) node()    {}
