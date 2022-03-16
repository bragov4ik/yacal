package builtin

import (
	"math"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/k0kubun/pp"
)

func UnaryOperation(i *types.Interpreter, args []interface{}) (a interface{}, err error) {
	if l := len(args); l != 1 {
		return nil, pp.Errorf("expected 1 arguments, but got %v", l)
	}
	args, err = i.EvalArgs(args)
	if err != nil {
		return
	}
	return args[0], nil

}

func BinaryOperation(i *types.Interpreter, args []interface{}) (a, b interface{}, err error) {
	if l := len(args); l != 2 {
		return nil, nil, pp.Errorf("expected 2 arguments, but got %v", l)
	}
	args, err = i.EvalArgs(args)
	if err != nil {
		return
	}
	return args[0], args[1], nil

}

func toFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return float64(v), nil
	case int:
		return float64(v), nil
	default:
		return math.NaN(), pp.Errorf("Expected number, but got %v", value)
	}
}

func toBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	default:
		return false, pp.Errorf("Expected bool, but got %v", value)
	}
}

func BinaryFloatOperation(i *types.Interpreter, args []interface{}) (a, b float64, err error) {
	_a, _b, err := BinaryOperation(i, args)
	if err != nil {
		return
	}
	a, err = toFloat64(_a)
	if err != nil {
		return
	}
	b, err = toFloat64(_b)
	return
}

func BinaryBoolOperation(i *types.Interpreter, args []interface{}) (a, b bool, err error) {
	_a, _b, err := BinaryOperation(i, args)
	if err != nil {
		return
	}
	a, err = toBool(_a)
	if err != nil {
		return
	}
	b, err = toBool(_b)
	return
}
