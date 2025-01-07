package luax

import (
	lua "github.com/yuin/gopher-lua"
)

// VM is an alias for lua.LState to represent the Lua virtual machine.
type VM = lua.LState

// Options is an alias for lua.Options to represent the options used when creating a Lua VM.
type Options = lua.Options

// NewVM creates a new Lua virtual machine (VM) instance.
func NewVM(opts ...Options) *VM {
	return lua.NewState(opts...)
}
