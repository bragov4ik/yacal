package builtin

import (
	"github.com/bragov4ik/yacal/pkg/interpreter/types"
)

func And(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryBoolOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a && b, nil
}

func Or(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryBoolOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a || b, nil
}

func Xor(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryBoolOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a != b, nil
}

func Not(i *types.Interpreter, args []interface{}) (interface{}, error) {
	_a, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	a, err := toBool(_a)
	if err != nil {
		return nil, err
	}
	return !a, nil
}
