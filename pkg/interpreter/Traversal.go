package interpreter

import "github.com/bragov4ik/yacal/pkg/parser/ast"

func ListTraversal(list ast.List) interface{} {
	var res interface{} = ast.Empty{}
	for i := len(list) - 1; i >= 0; i-- {
		res = ast.Cons{Val: ElemTraversal(list[i]), Next: res}
	}
	return res
}

func ElemTraversal(elem interface{}) interface{} {
	if list, ok := elem.(ast.List); ok {
		return ListTraversal(list)
	}
	return elem
}

func TreeTraversal(ast []interface{}) []interface{} {
	var res = make([]interface{}, len(ast))
	for i, elem := range ast {
		res[i] = ElemTraversal(elem)
	}
	return res
}
