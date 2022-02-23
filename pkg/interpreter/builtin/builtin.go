package builtin

import (
	"github.com/k0kubun/pp"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func plusInt(args []interface{}) (interface{}, error) {
	acc := int(0)
	for _, a := range args {
		if n, ok := a.(int); ok {
			acc += n
		} else {
			return nil, pp.Errorf("Expected number, but got %v", a)
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
			return nil, pp.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func Plus(in *types.Interpreter, args []interface{}) (interface{}, error) {
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
		return nil, pp.Errorf("Expected number, but got %v", n)
	} else {
		acc = n
	}

	for _, a := range args[1:] {
		if n, ok := a.(int); ok {
			acc -= n
		} else {
			return nil, pp.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func minusReal(args []interface{}) (interface{}, error) {
	acc := float64(0)
	if n, ok := args[0].(float64); !ok {
		return nil, pp.Errorf("Expected number, but got %v", n)
	} else {
		acc = n
	}

	for _, a := range args[1:] {
		if n, ok := a.(float64); ok {
			acc -= n
		} else {
			return nil, pp.Errorf("Expected number, but got %v", a)
		}
	}
	return acc, nil
}

func Minus(in *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, pp.Errorf("Expected at least 1 number")
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
