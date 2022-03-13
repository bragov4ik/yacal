package builtin

import (
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func Lambda(_ *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, pp.Errorf("Expected argument list and function body")
	}

	arglist := []string{}
	al, ok := args[0].(ast.List)
	if !ok {
		return nil, pp.Errorf("Expected argument list and function body")
	}
	for _, a := range al {
		if arg, ok := a.(ast.Atom); ok {
			arglist = append(arglist, arg.Val)
		} else {
			return nil, pp.Errorf("Expected argument list and function body")
		}
	}

	body := args[1]
	lambda := func(in *types.Interpreter, args []interface{}) (interface{}, error) {
		if len(arglist) != len(args) {
			return nil, pp.Errorf("Expected %v arguments for function, but have %v", len(arglist), len(args))
		}

		old_state := map[string]interface{}{}
		for i := 0; i < len(args); i++ {
			if old := in.SetState(arglist[i], args[i]); old != nil {
				old_state[arglist[i]] = old
			}
		}

		v, err := in.Eval(body)

		for k, v := range old_state {
			in.SetState(k, v)
		}

		return v, err
	}

	return types.Func(lambda), nil
}

func Set(i *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, pp.Errorf("Expected 2 argument for set")
	}
	arg, ok := args[0].(ast.Atom)
	if !ok {
		return nil, pp.Errorf("Expected Atom as first argument")
	}
	v, err := i.Eval(args[1])
	if err != nil {
		return nil, err
	}

	// Return old value
	return i.SetState(arg.Val, v), nil
}

func Quote(_ *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, pp.Errorf("Expected only 1 argument for quote")
	}
	// Do not evaluate argument
	return args[0], nil
}

func Eval(i *types.Interpreter, args []interface{}) (interface{}, error) {
	to_evaluate, err := UnaryOperation(i, args)
	if err != nil {
		return nil, err
	}
	return i.Eval(to_evaluate)
}

func Cond(i *types.Interpreter, args []interface{}) (interface{}, error) {
	if l := len(args); !(l == 2 || l == 3) {
		return nil, pp.Errorf("Expected 2 or 3 arguments for cond")
	}
	condition, err := i.Eval(args[0])
	if err != nil {
		return nil, err
	}
	if _, ok := condition.(bool); !ok {
		return nil, pp.Errorf("Expected bool in first argument of cond, but got %v", condition)
	}
	if condition.(bool) {
		success_statment, err := i.Eval(args[1])
		if err != nil {
			return nil, err
		}
		return success_statment, nil
	} else {
		if len(args) == 2 {
			return ast.Null{}, nil
		} else {
			failed_statment, err := i.Eval(args[2])
			if err != nil {
				return nil, err
			}
			return failed_statment, nil
		}
	}
}
