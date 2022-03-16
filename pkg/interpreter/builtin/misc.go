package builtin

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func Lambda(_ *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Expected argument list and function body")
	}

	arglist := []string{}
	al, ok := args[0].(ast.List)
	if !ok {
		return nil, fmt.Errorf("Expected argument list and function body")
	}
	for _, a := range al {
		if arg, ok := a.(ast.Atom); ok {
			arglist = append(arglist, arg.Val)
		} else {
			return nil, fmt.Errorf("Expected argument list and function body")
		}
	}

	body := args[1]
	lambda := func(in *types.Interpreter, args []interface{}) (interface{}, error) {
		if len(arglist) != len(args) {
			return nil, fmt.Errorf("Expected %v arguments for function, but have %v", len(arglist), len(args))
		}

		old_state := map[string]interface{}{}
		for i := 0; i < len(args); i++ {
			if old := in.SetState(arglist[i], args[i]); old != nil {
				old_state[arglist[i]] = old
			}
		}

		v, err := in.Eval(body)

		for k, v := range old_state {
			if v == nil {
				in.DeleteState(k)
			} else {
				in.SetState(k, v)
			}
		}

		return v, err
	}

	return types.Func(lambda), nil
}

func Set(i *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Expected 2 argument for set")
	}
	arg, ok := args[0].(ast.Atom)
	if !ok {
		return nil, fmt.Errorf("Expected Atom as first argument")
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
		return nil, fmt.Errorf("Expected only 1 argument for quote")
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
		return nil, fmt.Errorf("Expected 2 or 3 arguments for cond")
	}
	condition, err := i.Eval(args[0])
	if err != nil {
		return nil, err
	}
	if _, ok := condition.(bool); !ok {
		return nil, fmt.Errorf("Expected bool in first argument of cond, but got %v", condition)
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

func While(i *types.Interpreter, args []interface{}) (interface{}, error) {
	if l := len(args); l != 2 {
		return nil, fmt.Errorf("expected 2 arguments for while, but got %v", l)
	}
	condition_statment, body_statment := args[0], args[1]
	for iter_number := 0; iter_number < 10; iter_number++ {
		condition, err := i.Eval(condition_statment)
		if err != nil {
			return nil, err
		}
		if _, ok := condition.(bool); !ok {
			return nil, fmt.Errorf("expected bool in first argument of while, but got %v", condition)
		}
		if condition.(bool) {
			fmt.Println(i.GetState("a"))
			i.Eval(body_statment)
			fmt.Println(i.GetState("a"))

		} else {
			break
		}
	}
	return ast.Null{}, nil
}

func Prog(i *types.Interpreter, args []interface{}) (interface{}, error) {
	if l := len(args); l != 2 {
		return nil, fmt.Errorf("expected 2 arguments for prog, but got %v", l)
	}

	atoms_context := []string{}
	al, ok := args[0].(ast.List)
	if !ok {
		return nil, fmt.Errorf("Expected argument list and program body")
	}
	for _, a := range al {
		if arg, ok := a.(ast.Atom); ok {
			atoms_context = append(atoms_context, arg.Val)
		} else {
			return nil, fmt.Errorf("Expected atoms in argument list")
		}
	}
	body, ok := args[1].(ast.List)
	if !ok {
		return nil, fmt.Errorf("Expected program body")
	}

	// Save state
	old_state := map[string]interface{}{}
	for iter_number := 0; iter_number < len(args); iter_number++ {
		// Save nils to delete from context later
		old_state[atoms_context[iter_number]], _ = i.GetState(atoms_context[iter_number])
	}

	var res interface{} = ast.Null{}
	for _, st := range body {
		_res, err := i.Eval(st)
		res = _res
		if err != nil {
			return nil, err
		}
	}

	// Restore state
	for k, v := range old_state {
		if v == nil {
			i.DeleteState(k)
		} else {
			i.SetState(k, v)
		}
	}
	// return result of last statement in prog
	return res, nil
}

func ToString(arg interface{}) string {
	switch arg.(type) {
	case string:
		return "\"" + fmt.Sprint(arg) + "\""
	case int:
		return fmt.Sprint(arg)
	case bool:
		return fmt.Sprint(arg)
	case rune:
		return fmt.Sprint(arg)
	case float64:
		return fmt.Sprint(arg)
	case ast.Null:
		return "null"
	case ast.Atom:
		return fmt.Sprint(arg.(ast.Atom).Val)
	case ast.List:
		list := arg.(ast.List)
		if val, ok := list[0].(ast.Atom); ok && val.Val == "quote" {
			return ToString(list[1])
		}
		var ret []string
		for _, elem := range list {
			ret = append(ret, ToString(elem))
		}
		return "(" + strings.Join(ret, " ") + ")"
	default:
		return ""
	}
}

func Print(i *types.Interpreter, args []interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return nil, err
	}
	var output []string
	for _, arg := range args {
		output = append(output, ToString(arg))
	}
	fmt.Println(strings.Join(output, " "))
	return nil, nil
}

func Input(_ *types.Interpreter, args []interface{}) (interface{}, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expected 0 arguments for input")
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text[:len(text)-1], nil
}
