package builtin

import (
	"fmt"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
	"math"
)

func Len(arg interface{}) int {
	res := 0
	cur := arg
	for _, ok := cur.(ast.Empty); !ok; _, ok = cur.(ast.Empty) {
		res++
		cons, _ := cur.(ast.Cons)
		cur = cons.Next
	}
	return res
}

func UnaryOperation(args interface{}) (a interface{}, err error) {
	if l := Len(args); l != 1 {
		return nil, fmt.Errorf("expected 1 arguments, but got %v", l)
	}
	cons, _ := args.(ast.Cons)
	return cons.Val, nil

}

func BinaryOperation(args interface{}) (a, b interface{}, err error) {
	if l := Len(args); l != 2 {
		return nil, nil, fmt.Errorf("expected 2 arguments, but got %v", l)
	}
	cons1, _ := args.(ast.Cons)
	cons2, _ := cons1.Next.(ast.Cons)
	return cons1.Val, cons2.Val, nil
}
func TernaryOperation(args interface{}) (a, b, c interface{}, err error) {
	if l := Len(args); l != 3 {
		return nil, nil, nil, fmt.Errorf("expected 2 arguments, but got %v", l)
	}
	cons1, _ := args.(ast.Cons)
	cons2, _ := cons1.Next.(ast.Cons)
	cons3, _ := cons2.Next.(ast.Cons)
	return cons1.Val, cons2.Val, cons3.Val, nil
}

func toFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	default:
		return math.NaN(), fmt.Errorf("Expected number, but got %T", value)
	}
}

func toBool(value interface{}) (bool, error) {

	switch v := value.(type) {
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("Expected bool, but got %v", value)
	}
}

func BinaryFloatOperation(args interface{}) (a, b float64, err error) {
	_a, _b, err := BinaryOperation(args)
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

func BinaryBoolOperation(args interface{}) (a, b bool, err error) {
	_a, _b, err := BinaryOperation(args)
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
