package luax

import (
	lua "github.com/yuin/gopher-lua"
)

// LValue wraps a lua.LValue and returns a custom Value type.
func LValue(v lua.LValue) Value {
	if v == nil {
		v = lua.LNil
	}
	return Value{v: v}
}

// Value is a custom type that wraps a lua.LValue.
type Value struct {
	v lua.LValue
}

// String returns the string representation of the wrapped value if it's not nil, otherwise an empty string.
func (v Value) String() string {
	if !v.IsNil() {
		return v.v.String()
	}
	return ""
}

// Type returns the type of the wrapped lua.LValue.
func (v Value) Type() lua.LValueType {
	if !v.IsNil() {
		return v.v.Type()
	}
	return lua.LTNil
}

// IsNil checks if the wrapped lua.LValue is nil or represents lua.LNil.
func (v Value) IsNil() bool {
	return v.v == nil || v.v == lua.LNil
}

// ToString retrieves the string value from the wrapped lua.LValue if it's of type string.
func (v Value) ToString() string {
	if !v.IsNil() && v.v.Type() == lua.LTString {
		if s, ok := v.v.(lua.LString); ok {
			return string(s)
		}
	}
	return ""
}

// ToInt8 retrieves the int8 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToInt8() int8 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return int8(av)
		}
	}
	return 0
}

// ToInt16 retrieves the int16 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToInt16() int16 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return int16(av)
		}
	}
	return 0
}

// ToInt32 retrieves the int32 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToInt32() int32 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return int32(av)
		}
	}
	return 0
}

// ToInt retrieves the int value from the wrapped lua.LValue if it's of type number.
func (v Value) ToInt() int {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return int(av)
		}
	}
	return 0
}

// ToInt64 retrieves the int64 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToInt64() int64 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return int64(av)
		}
	}
	return 0
}

// ToUint8 retrieves the uint8 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToUint8() uint8 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return uint8(av)
		}
	}
	return 0
}

// ToUint16 retrieves the uint16 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToUint16() uint16 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return uint16(av)
		}
	}
	return 0
}

// ToUint32 retrieves the uint32 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToUint32() uint32 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return uint32(av)
		}
	}
	return 0
}

// ToUint retrieves the uint value from the wrapped lua.LValue if it's of type number.
func (v Value) ToUint() uint {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return uint(av)
		}
	}
	return 0
}

// ToUint64 retrieves the uint64 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToUint64() uint64 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return uint64(av)
		}
	}
	return 0
}

// ToBool retrieves the bool value from the wrapped lua.LValue if it's of type boolean.
func (v Value) ToBool() bool {
	if !v.IsNil() && v.v.Type() == lua.LTBool {
		if b, ok := v.v.(lua.LBool); ok {
			return bool(b)
		}
	}
	return false
}

// ToFloat32 retrieves the float32 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToFloat32() float32 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return float32(av)
		}
	}
	return 0.0
}

// ToFloat64 retrieves the float64 value from the wrapped lua.LValue if it's of type number.
func (v Value) ToFloat64() float64 {
	if !v.IsNil() && v.v.Type() == lua.LTNumber {
		if av, ok := v.v.(lua.LNumber); ok {
			return float64(av)
		}
	}
	return 0.0
}

// ToTable retrieves the table value from the wrapped lua.LValue if it's of type table.
func (v Value) ToTable() Table {
	if !v.IsNil() && v.v.Type() == lua.LTTable {
		if av, ok := v.v.(*lua.LTable); ok {
			return LTable(av)
		}
	}
	return LTable(nil)
}

// LValue returns the underlying lua.LValue.
func (v Value) LValue() lua.LValue {
	if v.IsNil() {
		return lua.LNil
	}
	return v.v
}

// IsNil checks if the lua.LValue is nil or represents lua.LNil.
func IsNil(v lua.LValue) bool {
	return LValue(v).IsNil()
}

// LValueToString retrieves the string value from a lua.LValue.
func LValueToString(v lua.LValue) string {
	return LValue(v).ToString()
}

// LValueToInt8 retrieves the int8 value from a lua.LValue.
func LValueToInt8(v lua.LValue) int8 {
	return LValue(v).ToInt8()
}

// LValueToInt16 retrieves the int16 value from a lua.LValue.
func LValueToInt16(v lua.LValue) int16 {
	return LValue(v).ToInt16()
}

// LValueToInt32 retrieves the int32 value from a lua.LValue.
func LValueToInt32(v lua.LValue) int32 {
	return LValue(v).ToInt32()
}

// LValueToInt retrieves the int value from a lua.LValue.
func LValueToInt(v lua.LValue) int {
	return LValue(v).ToInt()
}

// LValueToInt64 retrieves the int64 value from a lua.LValue.
func LValueToInt64(v lua.LValue) int64 {
	return LValue(v).ToInt64()
}

// LValueToUint8 retrieves the uint8 value from a lua.LValue.
func LValueToUint8(v lua.LValue) uint8 {
	return LValue(v).ToUint8()
}

// LValueToUint16 retrieves the uint16 value from a lua.LValue.
func LValueToUint16(v lua.LValue) uint16 {
	return LValue(v).ToUint16()
}

// LValueToUint32 retrieves the uint32 value from a lua.LValue.
func LValueToUint32(v lua.LValue) uint32 {
	return LValue(v).ToUint32()
}

// LValueToUint retrieves the uint value from a lua.LValue.
func LValueToUint(v lua.LValue) uint {
	return LValue(v).ToUint()
}

// LValueToUint64 retrieves the uint64 value from a lua.LValue.
func LValueToUint64(v lua.LValue) uint64 {
	return LValue(v).ToUint64()
}

// LValueToBool retrieves the bool value from a lua.LValue.
func LValueToBool(v lua.LValue) bool {
	return LValue(v).ToBool()
}

// LValueToFloat32 retrieves the float32 value from a lua.LValue.
func LValueToFloat32(v lua.LValue) float32 {
	return LValue(v).ToFloat32()
}

// LValueToFloat64 retrieves the float64 value from a lua.LValue.
func LValueToFloat64(v lua.LValue) float64 {
	return LValue(v).ToFloat64()
}

// LValueToArray retrieves the array of value from a lua.LValue
func LValueToArray(v lua.LValue) []Value {
	return LValue(v).ToTable().ToArray()
}

// LValueToTable retrieves the map of value from a lua.LValue.
func LValueToTable(v lua.LValue) map[string]Value {
	return LValue(v).ToTable().ToMap()
}

// LValueToDict retrieves the dict of value from a lua.LValue.
func LValueToDict(v lua.LValue) map[Value]Value {
	return LValue(v).ToTable().ToDict()
}
