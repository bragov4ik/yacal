package builtin

import (
	"fmt"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func Head(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	_l, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	l, ok := _l.(ast.Cons)
	if !ok {
		return nil, fmt.Errorf("expected non-empty list, but got %v", _l)
	}
	return l.Val, nil
}

func Tail(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	_l, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	l, ok := _l.(ast.Cons)
	if !ok {
		return nil, fmt.Errorf("expected non-empty list, but got %v", _l)
	}
	return l.Next, nil
}

func Cons(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	item, _l, err := BinaryOperation(args)
	if err != nil {
		return nil, err
	}
	if _, ok1 := _l.(ast.Cons); !ok1 {
		if _, ok2 := _l.(ast.Empty); !ok2 {
			return nil, fmt.Errorf("expected second argument to be list, but got %v", _l)
		}
	}
	return ast.Cons{Val: item, Next: _l}, nil
}
