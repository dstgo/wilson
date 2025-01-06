package luax

import lua "github.com/yuin/gopher-lua"

func LTable(t *lua.LTable) *Table {
	return (*Table)(t)
}

type Table lua.LTable

func (t *Table) Get(key string) lua.LValue {
	if (*lua.LTable)(t) == lua.LNil {
		return nil
	}
	return (*lua.LTable)(t).RawGetString(key)
}

func (t *Table) GetString(key string) string {
	lValue := t.Get(key)
	if lValue != nil && lValue.Type() == lua.LTString {
		return lValue.(lua.LString).String()
	}
	return ""
}

func (t *Table) GetInt64(key string) int64 {
	lValue := t.Get(key)
	if lValue != nil && lValue.Type() == lua.LTNumber {
		return int64(lValue.(lua.LNumber))
	}
	return 0
}

func (t *Table) GetFloat64(key string) float64 {
	lValue := t.Get(key)
	if lValue != nil && lValue.Type() == lua.LTNumber {
		return float64(lValue.(lua.LNumber))
	}
	return 0
}

func (t *Table) GetBool(key string) bool {
	lValue := t.Get(key)
	if lValue != nil && lValue.Type() == lua.LTBool {
		return bool(lValue.(lua.LBool))
	}
	return false
}

func (t *Table) GetTable(key string) *Table {
	lValue := t.Get(key)
	if lValue != nil && lValue.Type() == lua.LTTable {
		return (*Table)(lValue.(*lua.LTable))
	}
	return nil
}

func (t *Table) T() *lua.LTable {
	return (*lua.LTable)(t)
}

func JudgeOptionValue(value lua.LValue) any {
	switch value.Type() {
	default:
		return "unknown type"
	case lua.LTString:
		return lua.LVAsString(value)
	case lua.LTNumber:
		return float64(lua.LVAsNumber(value))
	case lua.LTBool:
		return lua.LVAsBool(value)
	}
}
