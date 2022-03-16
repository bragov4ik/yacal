package builtin

import (
	"fmt"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
)

func plusInt(args []interface{}) (interface{}, error) {
	acc := int(0)
	for _, a := range args {
		if n, ok := a.(int); ok {
			acc += n
		} else {
			return nil, fmt.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func plusReal(args []interface{}) (interface{}, error) {
	acc := float64(0.0)
	for _, a := range args {
		if n, ok := a.(float64); ok {
			acc += n
		} else {
			return nil, fmt.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func Plus(in *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("Expected at least 1 number")
	}
	args, err := in.EvalArgs(args)
	if err != nil {
		return nil, err
	}

	if v, err := plusInt(args); err == nil {
		return v, nil
	}
	return plusReal(args)
}

func minusInt(args []interface{}) (interface{}, error) {
	acc := int(0)
	if n, ok := args[0].(int); !ok {
		return nil, fmt.Errorf("Expected number, but got %v", n)
	} else {
		acc = n
	}

	for _, a := range args[1:] {
		if n, ok := a.(int); ok {
			acc -= n
		} else {
			return nil, fmt.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func minusReal(args []interface{}) (interface{}, error) {
	acc := float64(0)
	if n, ok := args[0].(float64); !ok {
		return nil, fmt.Errorf("Expected number, but got %v", n)
	} else {
		acc = n
	}

	for _, a := range args[1:] {
		if n, ok := a.(float64); ok {
			acc -= n
		} else {
			return nil, fmt.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func Minus(in *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("Expected at least 1 number")
	}
	args, err := in.EvalArgs(args)
	if err != nil {
		return nil, err
	}
	if v, err := minusInt(args); err == nil {
		return v, nil
	}
	return minusReal(args)
}

func timesInt(args []interface{}) (interface{}, error) {
	acc := int(1)
	for _, a := range args {
		if n, ok := a.(int); ok {
			acc *= n
		} else {
			return nil, fmt.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func timesReal(args []interface{}) (interface{}, error) {
	acc := float64(1.0)
	for _, a := range args {
		if n, ok := a.(float64); ok {
			acc *= n
		} else {
			return nil, fmt.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func Times(in *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("Expected at least 1 number")
	}
	args, err := in.EvalArgs(args)
	if err != nil {
		return nil, err
	}
	if v, err := timesInt(args); err == nil {
		return v, nil
	}
	return timesReal(args)
}

func divideInt(args []interface{}) (interface{}, error) {
	nom := int(1)
	if n, ok := args[0].(int); !ok {
		return nil, fmt.Errorf("Expected number, but got %v", n)
	} else {
		nom = n
	}
	denom := int(1)
	_denom, err := timesInt(args[1:])
	if err != nil {
		return nil, err
	}
	if n, ok := _denom.(int); !ok {
		return nil, fmt.Errorf("Expected number, but got %v", n)
	} else {
		denom = n
	}
	// Perform integer division
	return nom / denom, nil
}

func divideReal(args []interface{}) (interface{}, error) {
	nom := float64(1)
	if n, ok := args[0].(float64); !ok {
		return nil, fmt.Errorf("Expected number, but got %v", n)
	} else {
		nom = n
	}
	denom := float64(1)
	_denom, err := timesReal(args[1:])
	if err != nil {
		return nil, err
	}
	if n, ok := _denom.(float64); !ok {
		return nil, fmt.Errorf("Expected number, but got %v", n)
	} else {
		denom = n
	}
	return nom / denom, nil
}

func Divide(in *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Expected at least 2 numbers")
	}
	args, err := in.EvalArgs(args)
	if err != nil {
		return nil, err
	}
	if v, err := divideInt(args); err == nil {
		return v, nil
	}
	return divideReal(args)
}

func Quals(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryFloatOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a == b, nil
}

func NotQuals(i *types.Interpreter, args []interface{}) (interface{}, error) {
	equals, err := Quals(i, args)
	if err != nil {
		return nil, err
	}
	return !equals.(bool), nil
}

func Greater(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryFloatOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a > b, nil
}

func GreaterOrEq(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryFloatOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a >= b, nil
}

func Less(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryFloatOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a < b, nil
}

func LessOrEq(i *types.Interpreter, args []interface{}) (interface{}, error) {
	a, b, err := BinaryFloatOperation(i, args)
	if err != nil {
		return nil, err
	}
	return a <= b, nil
}
