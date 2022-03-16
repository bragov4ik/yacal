package builtin

import (
	"fmt"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func Head(i *types.Interpreter, args []interface{}) (interface{}, error) {
	_l, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	l, ok := _l.(ast.List)
	if !ok {
		return nil, fmt.Errorf("Expected list, but got %v", _l)
	}
	if len(l) == 0 {
		return ast.Null{}, nil
	}
	return l[0], nil
}

func Tail(i *types.Interpreter, args []interface{}) (interface{}, error) {
	_l, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	l, ok := _l.(ast.List)
	if !ok {
		return nil, fmt.Errorf("Expected list, but got %v", _l)
	}
	if len(l) < 2 {
		return ast.Null{}, nil
	}
	return l[1:], nil
}

func Cons(i *types.Interpreter, args []interface{}) (interface{}, error) {
	item, _l, err := BinaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	l, ok := _l.(ast.List)
	if !ok {
		return nil, fmt.Errorf("Expected list, but got %v", _l)
	}
	l = append(l, 0)
	copy(l[1:], l)
	l[0] = item
	return l, nil
}
