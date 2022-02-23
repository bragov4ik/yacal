package interpreter

import (
	"github.com/bragov4ik/yacal/pkg/interpreter/builtin"
	"github.com/bragov4ik/yacal/pkg/interpreter/types"
)

func initFunc() map[string]interface{} {
	return map[string]interface{}{
		"+":      types.Func(builtin.Plus),
		"-":      types.Func(builtin.Minus),
		"lambda": types.Func(builtin.Lambda),
		"quote":  types.Func(builtin.Quote),
		"set":    types.Func(builtin.Set),
	}
}

func New() *types.Interpreter { return types.NewInterpreter(initFunc()) }
