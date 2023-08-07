package luax

import (
	lua "github.com/yuin/gopher-lua"
)

func GetLuaValueFromTable(table *lua.LTable, key string) lua.LValue {
	return table.RawGet(lua.LString(key))
}

func GetStringFromTable(table *lua.LTable, key string) string {
	v := GetLuaValueFromTable(table, key)
	if v.Type() == lua.LTNil {
		return ""
	}
	return string(v.(lua.LString))
}
func GetNumberFromTable(table *lua.LTable, key string) float64 {
	v := GetLuaValueFromTable(table, key)
	if v.Type() == lua.LTNil {
		return 0
	}
	return float64(v.(lua.LNumber))
}
func GetBoolFromTable(table *lua.LTable, key string) bool {
	v := GetLuaValueFromTable(table, key)
	if v.Type() == lua.LTNil {
		return false
	}
	return bool(v.(lua.LBool))
}
func GetTableFromTable(table *lua.LTable, key string) *lua.LTable {
	v := GetLuaValueFromTable(table, key)
	if v.Type() == lua.LTNil {
		return new(lua.LTable)
	}
	return v.(*lua.LTable)
}
func GetChannelFromTable(table *lua.LTable, key string) lua.LChannel {
	v := GetLuaValueFromTable(table, key)
	if v.Type() == lua.LTNil {
		return make(lua.LChannel)
	}
	return v.(lua.LChannel)
}
