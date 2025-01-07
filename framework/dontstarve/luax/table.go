package luax

import lua "github.com/yuin/gopher-lua"

// Table is a wrapper type for lua.LTable
type Table struct {
	t *lua.LTable
}

// LTable converts a lua.LTable to a Table wrapper.
func LTable(t *lua.LTable) Table {
	if t == nil {
		t = &lua.LTable{}
	}
	return Table{t: t}
}

// Get retrieves a value from the table using a string key.
func (t Table) Get(key string) Value {
	if t.t != nil {
		return LValue(t.t.RawGetString(key))
	}
	return LValue(nil)
}

// Index retrieves a value from the table using an integer index.
func (t Table) Index(index int) Value {
	if t.t != nil {
		return LValue(t.t.RawGetInt(index))
	}
	return LValue(nil)
}

// GetString retrieves a string value from the table using a string key.
func (t Table) GetString(key string) string {
	return t.Get(key).ToString()
}

// GetInt8 retrieves an int8 value from the table using a string key.
func (t Table) GetInt8(key string) int8 {
	return t.Get(key).ToInt8()
}

// GetInt16 retrieves an int16 value from the table using a string key.
func (t Table) GetInt16(key string) int16 {
	return t.Get(key).ToInt16()
}

// GetInt32 retrieves an int32 value from the table using a string key.
func (t Table) GetInt32(key string) int32 {
	return t.Get(key).ToInt32()
}

// GetInt retrieves an int value from the table using a string key.
func (t Table) GetInt(key string) int {
	return t.Get(key).ToInt()
}

// GetInt64 retrieves an int64 value from the table using a string key.
func (t Table) GetInt64(key string) int64 {
	return t.Get(key).ToInt64()
}

// GetUint8 retrieves a uint8 value from the table using a string key.
func (t Table) GetUint8(key string) uint8 {
	return t.Get(key).ToUint8()
}

// GetUint16 retrieves a uint16 value from the table using a string key.
func (t Table) GetUint16(key string) uint16 {
	return t.Get(key).ToUint16()
}

// GetUint32 retrieves a uint32 value from the table using a string key.
func (t Table) GetUint32(key string) uint32 {
	return t.Get(key).ToUint32()
}

// GetUint retrieves a uint value from the table using a string key.
func (t Table) GetUint(key string) uint {
	return t.Get(key).ToUint()
}

// GetUint64 retrieves a uint64 value from the table using a string key.
func (t Table) GetUint64(key string) uint64 {
	return t.Get(key).ToUint64()
}

// GetFloat32 retrieves a float32 value from the table using a string key.
func (t Table) GetFloat32(key string) float32 {
	return t.Get(key).ToFloat32()
}

// GetFloat64 retrieves a float64 value from the table using a string key.
func (t Table) GetFloat64(key string) float64 {
	return t.Get(key).ToFloat64()
}

// GetBool retrieves a bool value from the table using a string key.
func (t Table) GetBool(key string) bool {
	return t.Get(key).ToBool()
}

// GetTable retrieves a nested table from the table using a string key.
func (t Table) GetTable(key string) Table {
	return t.Get(key).ToTable()
}

// LTable converts the Table back to a *lua.LTable.
func (t Table) LTable() *lua.LTable {
	if t.t != nil {
		return t.t
	}
	return &lua.LTable{}
}

// ArrayLen returns the length of the table array.
func (t Table) ArrayLen() int {
	return t.LTable().Len()
}

// MapLen returns the length of the table map.
// due to the limitation of Lua virtual machine, we can only traverse to get the length
func (t Table) MapLen() int {
	var length int
	t.LTable().ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() == lua.LTString {
			length++
		}
	})
	return length
}

// DictLen returns the length of the table dict.
// due to the limitation of Lua virtual machine, we can only traverse to get the length
func (t Table) DictLen() int {
	var length int
	t.LTable().ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() != lua.LTString && k.Type() != lua.LTNumber {
			length++
		}
	})
	return length
}

// ForEach iterates over the table using a function.
func (t Table) ForEach(fn func(k, v Value)) {
	t.LTable().ForEach(func(k lua.LValue, v lua.LValue) {
		fn(LValue(k), LValue(v))
	})
}

// ArrayForEach iterates over the table array using a function.
func (t Table) ArrayForEach(fn func(k int, v Value)) {
	t.LTable().ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() == lua.LTNumber {
			fn(LValueToInt(k), LValue(v))
		}
	})
}

// MapForEach iterates over the table map using a function.
func (t Table) MapForEach(fn func(k string, v Value)) {
	t.LTable().ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() == lua.LTString {
			fn(LValueToString(k), LValue(v))
		}
	})
}

// DictForEach iterates over the table dict using a function.
func (t Table) DictForEach(fn func(k, v Value)) {
	t.LTable().ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() != lua.LTString && k.Type() != lua.LTNumber {
			fn(LValue(k), LValue(v))
		}
	})
}

// ToArray converts the table into a slice of Value.
func (t Table) ToArray() []Value {
	var arr []Value
	t.LTable().ForEach(func(index lua.LValue, value lua.LValue) {
		if index.Type() == lua.LTNumber {
			arr = append(arr, LValue(value))
		}
	})
	return arr
}

// ToMap converts the table into a map of string to Value.
func (t Table) ToMap() map[string]Value {
	m := make(map[string]Value)
	t.LTable().ForEach(func(key lua.LValue, value lua.LValue) {
		if key.Type() == lua.LTString {
			m[LValueToString(key)] = LValue(value)
		}
	})
	return m
}

// ToDict converts the table into a map of Value to Value.
func (t Table) ToDict() map[Value]Value {
	m := make(map[Value]Value)
	t.LTable().ForEach(func(index lua.LValue, value lua.LValue) {
		if index.Type() != lua.LTNumber && index.Type() != lua.LTString {
			m[LValue(index)] = LValue(value)
		}
	})
	return m
}

// Next retrieves the next key-value pair in the table after a given key.
func (t Table) Next(key Value) (Value, Value) {
	nextKey, curVal := t.LTable().Next(key)
	return LValue(nextKey), LValue(curVal)
}

// LTableGet is a convenience function to get a value from a lua.LTable using a string key.
func LTableGet(t *lua.LTable, key string) Value {
	return LTable(t).Get(key)
}

// LTableIndex is a convenience function to get a value from a lua.LTable using an int key.
func LTableIndex(t *lua.LTable, i int) Value {
	return LTable(t).Index(i)
}

// LTableGetString is a convenience function to get a string value from a lua.LTable using a string key.
func LTableGetString(t *lua.LTable, key string) string {
	return LTable(t).GetString(key)
}

// LTableGetInt64 is a convenience function to get an int64 value from a lua.LTable using a string key.
func LTableGetInt64(t *lua.LTable, key string) int64 {
	return LTable(t).GetInt64(key)
}

// LTableGetInt is a convenience function to get an int value from a lua.LTable using a string key.
func LTableGetInt(t *lua.LTable, key string) int {
	return LTable(t).GetInt(key)
}

// LTableGetInt8 is a convenience function to get an int8 value from a lua.LTable using a string key.
func LTableGetInt8(t *lua.LTable, key string) int8 {
	return LTable(t).GetInt8(key)
}

// LTableGetInt16 is a convenience function to get an int16 value from a lua.LTable using a string key.
func LTableGetInt16(t *lua.LTable, key string) int16 {
	return LTable(t).GetInt16(key)
}

// LTableGetInt32 is a convenience function to get an int32 value from a lua.LTable using a string key.
func LTableGetInt32(t *lua.LTable, key string) int32 {
	return LTable(t).GetInt32(key)
}

// LTableGetUint is a convenience function to get a uint value from a lua.LTable using a string key.
func LTableGetUint(t *lua.LTable, key string) uint {
	return LTable(t).GetUint(key)
}

// LTableGetUint8 is a convenience function to get a uint8 value from a lua.LTable using a string key.
func LTableGetUint8(t *lua.LTable, key string) uint8 {
	return LTable(t).GetUint8(key)
}

// LTableGetUint16 is a convenience function to get a uint16 value from a lua.LTable using a string key.
func LTableGetUint16(t *lua.LTable, key string) uint16 {
	return LTable(t).GetUint16(key)
}

// LTableGetUint32 is a convenience function to get a uint32 value from a lua.LTable using a string key.
func LTableGetUint32(t *lua.LTable, key string) uint32 {
	return LTable(t).GetUint32(key)
}

// LTableGetUint64 is a convenience function to get a uint64 value from a lua.LTable using a string key.
func LTableGetUint64(t *lua.LTable, key string) uint64 {
	return LTable(t).GetUint64(key)
}

// LTableGetFloat64 is a convenience function to get a float64 value from a lua.LTable using a string key.
func LTableGetFloat64(t *lua.LTable, key string) float64 {
	return LTable(t).GetFloat64(key)
}

// LTableGetFloat32 is a convenience function to get a float32 value from a lua.LTable using a string key.
func LTableGetFloat32(t *lua.LTable, key string) float32 {
	return LTable(t).GetFloat32(key)
}

// LTableGetBool is a convenience function to get a bool value from a lua.LTable using a string key.
func LTableGetBool(t *lua.LTable, key string) bool {
	return LTable(t).GetBool(key)
}

// LTableGetTable is a convenience function to get a table value from a lua.LTable using a string key.
func LTableGetTable(t *lua.LTable, key string) Table {
	return LTable(t).GetTable(key)
}

// LTableToArray converts the lua.LTable to a slice of Value.
func LTableToArray(t *lua.LTable) []Value {
	return LTable(t).ToArray()
}

// LTableToMap converts the lua.LTable to a map of string to Value.
func LTableToMap(t *lua.LTable) map[string]Value {
	return LTable(t).ToMap()
}

// LTableToDict converts the lua.LTable to a map of Value to Value.
func LTableToDict(t *lua.LTable) map[Value]Value {
	return LTable(t).ToDict()
}
