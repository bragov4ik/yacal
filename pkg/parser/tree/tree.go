package tree

type Atom struct {
	Ident string
}

type Literal struct {
	Value interface{}
}

type List struct {
	Elements []interface{}
}
