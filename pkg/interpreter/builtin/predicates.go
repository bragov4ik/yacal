package builtin

import (
	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func IsInt(i *types.Interpreter, args []interface{}) (interface{}, error) {
	v, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	_, ok := v.(int)
	return ok, nil
}

func IsReal(i *types.Interpreter, args []interface{}) (interface{}, error) {
	v, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	_, ok := v.(float64)
	return ok, nil
}

func IsBool(i *types.Interpreter, args []interface{}) (interface{}, error) {
	v, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	_, ok := v.(bool)
	return ok, nil
}

func IsNull(i *types.Interpreter, args []interface{}) (interface{}, error) {
	v, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	_, ok := v.(ast.Null)
	return ok, nil
}

func IsAtom(i *types.Interpreter, args []interface{}) (interface{}, error) {
	v, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	_, ok := v.(ast.Atom)
	return ok, nil
}

func IsList(i *types.Interpreter, args []interface{}) (interface{}, error) {
	v, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	_, ok := v.(ast.List)
	return ok, nil
}
