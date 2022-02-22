package types

import (
	"fmt"

	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

type (
	Interpreter struct {
		state map[string]interface{}
	}

	Func func(i *Interpreter, args []interface{}) (interface{}, error)
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
	case rune:
		return expr, nil
	case float64:
		return expr, nil
	case Func:
		return expr, nil

	case ast.Atom:
		name := expr.(ast.Atom).Val
		if v, ok := i.state[name]; ok {
			return v, nil
		} else {
			return nil, pp.Errorf("Unexpected atom `%v'", name)
		}
	case ast.List:
		list := expr.(ast.List)
		if len(list) == 0 {
			return expr, nil
		}
		a, err := i.Eval(list[0])
		if err != nil {
			return nil, pp.Errorf("Error while evaluation of function (%v): %v", list[0], err)
		}
		f, ok := a.(Func)
		if !ok {
			return nil, pp.Errorf("Expected function, but got %v", f)
		}
		out, err := f(i, list[1:])
		if err != nil {
			return nil, fmt.Errorf("Error while executing function: %v", err)
		}
		return out, nil
	default:
		return nil, pp.Errorf("%v expression has unknown type", expr)
	}
}

func (in *Interpreter) EvalArgs(args []interface{}) ([]interface{}, error) {
	for i := 0; i < len(args); i++ {
		var err error
		if args[i], err = in.Eval(args[i]); err != nil {
			return nil, err
		}
	}
	return args, nil
}
