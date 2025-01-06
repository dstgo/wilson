package luax

import (
	lua "github.com/yuin/gopher-lua"
)

type VM = lua.LState

type Options = lua.Options

func NewVM(opts ...Options) *VM {
	return lua.NewState(opts...)
}
