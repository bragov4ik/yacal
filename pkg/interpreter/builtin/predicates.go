package builtin

import (
	"fmt"
	"strconv"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func IsInt(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	_, ok := v.(int)
	return ok, nil
}

func IsReal(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	_, ok := v.(float64)
	return ok, nil
}

func IsBool(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	_, ok := v.(bool)
	return ok, nil
}

func IsNull(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	_, ok := v.(ast.Null)
	return ok, nil
}

func IsAtom(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	_, ok := v.(ast.Atom)
	return ok, nil
}

func IsList(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	_, ok1 := v.(ast.Cons)
	_, ok2 := v.(ast.Empty)
	return ok1 || ok2, nil
}

func IsEmpty(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	_, ok := v.(ast.Empty)
	return ok, nil
}

func ToInt(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	string_v, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("expected string as first argument of toint, but got %v", v)
	}
	return strconv.Atoi(string_v)
}

func ToReal(i *types.Interpreter, args interface{}) (interface{}, error) {
	v, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	v, err = i.Eval(v)
	if err != nil {
		return nil, err
	}
	string_v, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("expected string as first argument of toreal, but got %v", v)
	}
	return strconv.ParseFloat(string_v, 64)
}
