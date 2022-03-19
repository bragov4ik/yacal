package builtin

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bragov4ik/yacal/pkg/interpreter/types"
	"github.com/bragov4ik/yacal/pkg/parser/ast"
)

func Lambda(_ *types.Interpreter, args interface{}) (interface{}, error) {
	arg1, body, err := BinaryOperation(args)
	if err != nil {
		return nil, fmt.Errorf("expected 2 arguments for lambda")
	}

	var arglist []string

	for _, isEmpty := arg1.(ast.Empty); !isEmpty; _, isEmpty = arg1.(ast.Empty) {
		cons, isCons := arg1.(ast.Cons)
		if !isCons {
			return nil, fmt.Errorf("exepcted first argument to be list")
		}
		arg, isAtom := cons.Val.(ast.Atom)
		if !isAtom {
			return nil, fmt.Errorf("expected first argument to contains only identifiers")
		}
		arglist = append(arglist, arg.Val)
		arg1 = cons.Next
	}

	lambda := func(in *types.Interpreter, args interface{}) (interface{}, error) {
		if l1, l2 := len(arglist), Len(args); l1 != l2 {
			return nil, fmt.Errorf("expected %v arguments for function, but have %v", l1, l2)
		}

		oldState := map[string]interface{}{}

		for i := 0; i < len(arglist); i++ {
			cons, _ := args.(ast.Cons)
			val, err := in.Eval(cons.Val)
			if err != nil {
				return nil, err
			}
			old := in.SetState(arglist[i], val)
			oldState[arglist[i]] = old
			args = cons.Next
		}

		v, err := in.Eval(body)

		for k, v := range oldState {
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

func Set(i *types.Interpreter, args interface{}) (interface{}, error) {
	arg1, arg2, err := BinaryOperation(args)
	arg, ok := arg1.(ast.Atom)
	if !ok {
		return nil, fmt.Errorf("expected Atom as first argument")
	}
	v, err := i.Eval(arg2)
	if err != nil {
		return nil, err
	}

	// Return old value
	return i.SetState(arg.Val, v), nil
}

func SetFunc(i *types.Interpreter, args interface{}) (interface{}, error) {
	arg1, arg2, arg3, err := TernaryOperation(args)

	name, ok := arg1.(ast.Atom)
	if !ok {
		return nil, fmt.Errorf("expected Atom as first argument")
	}

	lambda, err := Lambda(i, ast.Cons{Val: arg2, Next: ast.Cons{Val: arg3, Next: ast.Empty{}}})
	if err != nil {
		return nil, err
	}

	return i.SetState(name.Val, lambda), nil
}

func Quote(_ *types.Interpreter, args interface{}) (interface{}, error) {
	res, err := UnaryOperation(args)

	if err != nil {
		return nil, err
	}

	// Do not evaluate argument
	return res, nil
}

func Eval(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return args, err
	}
	toEvaluate, err := UnaryOperation(args)
	if err != nil {
		return nil, err
	}
	return i.Eval(toEvaluate)
}

func Cond(i *types.Interpreter, args interface{}) (interface{}, error) {
	if l := Len(args); !(l == 2 || l == 3) {
		return nil, fmt.Errorf("expected 2 or 3 arguments for cond")
	}
	arg1, _ := args.(ast.Cons)
	arg2, _ := arg1.Next.(ast.Cons)
	condition, err := i.Eval(arg1.Val)
	if err != nil {
		return nil, err
	}
	if _, ok := condition.(bool); !ok {
		return nil, fmt.Errorf("expected bool in first argument of cond, but got %v", condition)
	}
	if condition.(bool) {
		successStatment, err := i.Eval(arg2.Val)
		if err != nil {
			return nil, err
		}
		return successStatment, nil
	} else {
		if arg3, ok := arg2.Next.(ast.Cons); !ok {
			return ast.Null{}, nil
		} else {
			failedStatment, err := i.Eval(arg3.Val)
			if err != nil {
				return nil, err
			}
			return failedStatment, nil
		}
	}
}

func While(i *types.Interpreter, args interface{}) (interface{}, error) {
	conditionStatment, bodyStatment, err := BinaryOperation(args)
	if err != nil {
		return nil, err
	}
	for {
		condition, err := i.Eval(conditionStatment)
		if err != nil {
			return nil, err
		}
		if _, ok := condition.(bool); !ok {
			return nil, fmt.Errorf("expected bool in first argument of while, but got %v", condition)
		}
		if condition.(bool) {
			_, err = i.Eval(bodyStatment)
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}
	return ast.Null{}, nil
}

func Prog(i *types.Interpreter, args interface{}) (interface{}, error) {
	arg1, body, err := BinaryOperation(args)
	if err != nil {
		return nil, err
	}

	var atomsContext []string
	for _, isEmpty := arg1.(ast.Empty); !isEmpty; _, isEmpty = arg1.(ast.Empty) {
		cons, isCons := arg1.(ast.Cons)
		if !isCons {
			return nil, fmt.Errorf("exepcted first argument to be list")
		}
		arg, isAtom := cons.Val.(ast.Atom)
		if !isAtom {
			return nil, fmt.Errorf("expected first argument to contains only identifiers")
		}
		atomsContext = append(atomsContext, arg.Val)
		arg1 = cons.Next
	}

	// Save state
	oldState := map[string]interface{}{}
	for _, a := range atomsContext {
		// Save nils to delete from context later
		oldState[a], _ = i.GetState(a)
	}

	var res interface{} = ast.Null{}
	for _, isEmpty := body.(ast.Empty); !isEmpty; _, isEmpty = body.(ast.Empty) {
		cons, isCons := body.(ast.Cons)
		if !isCons {
			return nil, fmt.Errorf("expected body to be list")
		}
		res, err = i.Eval(cons.Val)
		if err != nil {
			return nil, err
		}
		body = cons.Next
	}

	// Restore state
	for k, v := range oldState {
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
	switch v := arg.(type) {
	case string:
		return "\"" + fmt.Sprint(arg) + "\""
	case int:
		return fmt.Sprint(arg)
	case bool:
		return fmt.Sprint(arg)
	case float64:
		return fmt.Sprint(arg)
	case types.Func:
		return fmt.Sprintf("{func: %s}", arg)
	case ast.Null:
		return "null"
	case ast.Atom:
		return fmt.Sprint(v.Val)
	case ast.Empty:
		return "()"
	case ast.Cons:
		var ret []string
		var list interface{} = v
		for _, isEmpty := list.(ast.Empty); !isEmpty; _, isEmpty = list.(ast.Empty) {
			cons, _ := list.(ast.Cons)
			ret = append(ret, ToString(cons.Val))
			list = cons.Next
		}
		return "(" + strings.Join(ret, " ") + ")"
	default:
		return ""
	}
}

func Print(i *types.Interpreter, args interface{}) (interface{}, error) {
	args, err := i.EvalArgs(args)
	if err != nil {
		return nil, err
	}
	var output []string
	for _, isEmpty := args.(ast.Empty); !isEmpty; _, isEmpty = args.(ast.Empty) {
		cons, _ := args.(ast.Cons)
		output = append(output, ToString(cons.Val))
		args = cons.Next
	}
	fmt.Println(strings.Join(output, " "))
	return ast.Null{}, nil
}

func Input(_ *types.Interpreter, args interface{}) (interface{}, error) {
	if Len(args) > 0 {
		return nil, fmt.Errorf("expected 0 arguments for input")
	}
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return strings.Trim(text, "\r\n\t "), nil
}
