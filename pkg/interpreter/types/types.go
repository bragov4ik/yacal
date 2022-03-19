package types

import (
	"fmt"

	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

type (
	Interpreter struct {
		state map[string]interface{}
	}

	Func func(i *Interpreter, args interface{}) (interface{}, error)
)

func NewInterpreter(state map[string]interface{}) *Interpreter { return &Interpreter{state} }

func (i *Interpreter) SetState(name string, val interface{}) interface{} {
	ret := interface{}(nil)
	if v, ok := i.state[name]; ok {
		ret = v
	}
	i.state[name] = val
	return ret
}

func (i *Interpreter) DeleteState(name string) {
	delete(i.state, name)
}

func (i *Interpreter) GetState(name string) (interface{}, bool) {
	v, ok := i.state[name]
	return v, ok
}

func (i *Interpreter) Eval(expr interface{}) (interface{}, error) {
	switch expr.(type) {
	case string:
		return expr, nil
	case int:
		return expr, nil
	case bool:
		return expr, nil
	case ast.Null:
		return expr, nil
	case float64:
		return expr, nil
	case Func:
		return expr, nil
	case ast.Empty:
		return expr, nil

	case ast.Atom:
		name := expr.(ast.Atom).Val
		if v, ok := i.state[name]; ok {
			return v, nil
		} else {
			return nil, fmt.Errorf("Unexpected atom '%v'", name)
		}
	case ast.Cons:
		list := expr.(ast.Cons)
		a, err := i.Eval(list.Val)
		if err != nil {
			return nil, fmt.Errorf("Error while evaluation of function (%v): %v", list.Val, err)
		}
		f, ok := a.(Func)
		if !ok {
			return nil, fmt.Errorf("Expected function, but got %v", f)
		}
		out, err := f(i, list.Next)
		if err != nil {
			return nil, fmt.Errorf("Error while executing function: %v", err)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("%v expression has unknown type", expr)
	}
}

func (i *Interpreter) EvalArgs(args interface{}) (interface{}, error) {
	switch arg := args.(type) {
	case ast.Cons:
		var err error
		arg.Val, err = i.Eval(arg.Val)
		if err != nil {
			return nil, err
		}
		arg.Next, err = i.EvalArgs(arg.Next)
		if err != nil {
			return nil, err
		}
		return arg, nil
	case ast.Empty:
		return arg, nil
	default:
		return nil, fmt.Errorf("unexpected type %T", arg)
	}
}
