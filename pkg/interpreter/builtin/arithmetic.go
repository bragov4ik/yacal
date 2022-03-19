package builtin

import (
	"fmt"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
)

func Plus(in *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := in.EvalArgs(args)
	if err != nil {
		return args, err
	}
	arg1, arg2, err := BinaryOperation(args)
	int1, ok1 := arg1.(int)
	int2, ok2 := arg2.(int)
	if ok1 && ok2 {
		return int1 + int2, nil
	}
	float1, ok1 := arg1.(float64)
	float2, ok2 := arg2.(float64)
	if ok1 && ok2 {
		return float1 + float2, nil
	}
	return nil, fmt.Errorf("expected numbers, but got %v %v", arg1, arg2)
}

func Minus(in *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := in.EvalArgs(args)
	if err != nil {
		return args, err
	}
	arg1, arg2, err := BinaryOperation(args)
	int1, ok1 := arg1.(int)
	int2, ok2 := arg2.(int)
	if ok1 && ok2 {
		return int1 - int2, nil
	}
	float1, ok1 := arg1.(float64)
	float2, ok2 := arg2.(float64)
	if ok1 && ok2 {
		return float1 - float2, nil
	}
	return nil, fmt.Errorf("expected numbers, but got %v %v", arg1, arg2)
}

func Times(in *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := in.EvalArgs(args)
	if err != nil {
		return args, err
	}
	arg1, arg2, err := BinaryOperation(args)
	int1, ok1 := arg1.(int)
	int2, ok2 := arg2.(int)
	if ok1 && ok2 {
		return int1 * int2, nil
	}
	float1, ok1 := arg1.(float64)
	float2, ok2 := arg2.(float64)
	if ok1 && ok2 {
		return float1 * float2, nil
	}
	return nil, fmt.Errorf("expected numbers, but got %v %v", arg1, arg2)
}

func Divide(in *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := in.EvalArgs(args)
	if err != nil {
		return args, err
	}
	arg1, arg2, err := BinaryOperation(args)
	int1, ok1 := arg1.(int)
	int2, ok2 := arg2.(int)
	if ok1 && ok2 {
		return int1 / int2, nil
	}
	float1, ok1 := arg1.(float64)
	float2, ok2 := arg2.(float64)
	if ok1 && ok2 {
		return float1 / float2, nil
	}
	return nil, fmt.Errorf("expected numbers, but got %v %v", arg1, arg2)
}

func Quals(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	a, b, err := BinaryFloatOperation(args)
	if err != nil {
		return nil, err
	}
	return a == b, nil
}

func NotQuals(i *types.Interpreter, args interface{}) (interface{}, error) {
	equals, err := Quals(i, args)
	if err != nil {
		return nil, err
	}
	return !equals.(bool), nil
}

func Greater(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	a, b, err := BinaryFloatOperation(args)
	if err != nil {
		return nil, err
	}
	return a > b, nil
}

func GreaterOrEq(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	a, b, err := BinaryFloatOperation(args)
	if err != nil {
		return nil, err
	}
	return a >= b, nil
}

func Less(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	a, b, err := BinaryFloatOperation(args)
	if err != nil {
		return nil, err
	}
	return a < b, nil
}

func LessOrEq(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	a, b, err := BinaryFloatOperation(args)
	if err != nil {
		return nil, err
	}
	return a <= b, nil
}
