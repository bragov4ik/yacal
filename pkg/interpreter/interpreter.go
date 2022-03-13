package interpreter

import (
	"github.com/bragov4ik/yacal/pkg/interpreter/builtin"
	"github.com/bragov4ik/yacal/pkg/interpreter/types"
)

func initFunc() map[string]interface{} {
	return map[string]interface{}{
		// TODO: times, minus, head, tail, cons, prog, while, return
		"+":      types.Func(builtin.Plus),
		"-":      types.Func(builtin.Minus),
		"cond":   types.Func(builtin.Cond),
		"lambda": types.Func(builtin.Lambda),
		"quote":  types.Func(builtin.Quote),
		"set":    types.Func(builtin.Set),
		"eval":   types.Func(builtin.Eval),
		"=":      types.Func(builtin.Quals),
		"!=":     types.Func(builtin.NotQuals),
		">":      types.Func(builtin.Greater),
		">=":     types.Func(builtin.GreaterOrEq),
		"<":      types.Func(builtin.Less),
		"<=":     types.Func(builtin.LessOrEq),
		"isreal": types.Func(builtin.IsReal),
		"isbool": types.Func(builtin.IsBool),
		"isnull": types.Func(builtin.IsNull),
		"isatom": types.Func(builtin.IsAtom),
		"islist": types.Func(builtin.IsList),
		"and":    types.Func(builtin.And),
		"or":     types.Func(builtin.Or),
		"xor":    types.Func(builtin.Xor),
		"not":    types.Func(builtin.Not),
	}
}

func New() *types.Interpreter { return types.NewInterpreter(initFunc()) }
