package interpreter

import (
	"github.com/bragov4ik/yacal/pkg/interpreter/builtin"
	"github.com/bragov4ik/yacal/pkg/interpreter/types"
)

func initFunc() map[string]interface{} {
	return map[string]interface{}{
		"+":       types.Func(builtin.Plus),
		"-":       types.Func(builtin.Minus),
		"*":       types.Func(builtin.Times),
		"/":       types.Func(builtin.Divide),
		"cond":    types.Func(builtin.Cond),
		"lambda":  types.Func(builtin.Lambda),
		"func":    types.Func(builtin.SetFunc),
		"quote":   types.Func(builtin.Quote),
		"set":     types.Func(builtin.Set),
		"eval":    types.Func(builtin.Eval),
		"while":   types.Func(builtin.While),
		"prog":    types.Func(builtin.Prog),
		"=":       types.Func(builtin.Quals),
		"!=":      types.Func(builtin.NotQuals),
		">":       types.Func(builtin.Greater),
		">=":      types.Func(builtin.GreaterOrEq),
		"<":       types.Func(builtin.Less),
		"<=":      types.Func(builtin.LessOrEq),
		"toint":   types.Func(builtin.ToInt),
		"isint":   types.Func(builtin.IsInt),
		"toreal":  types.Func(builtin.ToReal),
		"isreal":  types.Func(builtin.IsReal),
		"isbool":  types.Func(builtin.IsBool),
		"isnull":  types.Func(builtin.IsNull),
		"isatom":  types.Func(builtin.IsAtom),
		"islist":  types.Func(builtin.IsList),
		"isempty": types.Func(builtin.IsEmpty),
		"and":     types.Func(builtin.And),
		"or":      types.Func(builtin.Or),
		"xor":     types.Func(builtin.Xor),
		"not":     types.Func(builtin.Not),
		"print":   types.Func(builtin.Print),
		"input":   types.Func(builtin.Input),
		"head":    types.Func(builtin.Head),
		"tail":    types.Func(builtin.Tail),
		"cons":    types.Func(builtin.Cons),
	}
}

func New() *types.Interpreter { return types.NewInterpreter(initFunc()) }
