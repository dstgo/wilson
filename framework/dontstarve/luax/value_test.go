package luax

import (
	"testing"

	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

func TestIsNil(t *testing.T) {
	// Test when the value is nil
	assert.True(t, IsNil(nil))      // nil should be considered Nil
	assert.True(t, IsNil(lua.LNil)) // lua.LNil should be considered Nil

	// Test when the value is not nil
	assert.False(t, IsNil(lua.LNumber(1)))      // lua.LNumber is not nil
	assert.False(t, IsNil(lua.LString("test"))) // lua.LString is not nil
	assert.False(t, IsNil(lua.LBool(true)))     // lua.LBool is not nil
	assert.False(t, IsNil(&lua.LTable{}))       // lua.LTable is not nil
}

func TestLValueType(t *testing.T) {
	// Test types of different Lua values
	assert.Equal(t, lua.LTNil, LValue(nil).Type())
	assert.Equal(t, lua.LTNil, LValue(lua.LNil).Type())
	assert.Equal(t, lua.LTNumber, LValue(lua.LNumber(10)).Type())
	assert.Equal(t, lua.LTString, LValue(lua.LString("hello")).Type())
	assert.Equal(t, lua.LTBool, LValue(lua.LBool(true)).Type())
	assert.Equal(t, lua.LTTable, LValue(&lua.LTable{}).Type())
}

func TestLValueToString(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, "", LValue(nil).ToString())
	assert.Equal(t, "", LValue(lua.LNil).ToString())

	// Test when the value is a string
	assert.Equal(t, "hello", LValue(lua.LString("hello")).ToString())

	// Test when the value is a number (should return empty string)
	assert.Equal(t, "", LValue(lua.LNumber(1)).ToString())
}

func TestLValueToInt8(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, int8(0), LValue(nil).ToInt8())
	assert.Equal(t, int8(0), LValue(lua.LNil).ToInt8())

	// Test when the value is a number
	assert.Equal(t, int8(1), LValue(lua.LNumber(1)).ToInt8())

	// Test when the value is a string (should return 0)
	assert.Equal(t, int8(0), LValue(lua.LString("test")).ToInt8())
}

func TestLValueToInt16(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, int16(0), LValue(nil).ToInt16())
	assert.Equal(t, int16(0), LValue(lua.LNil).ToInt16())

	// Test when the value is a number
	assert.Equal(t, int16(1), LValue(lua.LNumber(1)).ToInt16())

	// Test when the value is a string (should return 0)
	assert.Equal(t, int16(0), LValue(lua.LString("test")).ToInt16())
}

func TestLValueToInt32(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, int32(0), LValue(nil).ToInt32())
	assert.Equal(t, int32(0), LValue(lua.LNil).ToInt32())

	// Test when the value is a number
	assert.Equal(t, int32(1), LValue(lua.LNumber(1)).ToInt32())

	// Test when the value is a string (should return 0)
	assert.Equal(t, int32(0), LValue(lua.LString("test")).ToInt32())
}

func TestLValueToInt64(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, int64(0), LValue(nil).ToInt64())
	assert.Equal(t, int64(0), LValue(lua.LNil).ToInt64())

	// Test when the value is a number
	assert.Equal(t, int64(1), LValue(lua.LNumber(1)).ToInt64())

	// Test when the value is a string (should return 0)
	assert.Equal(t, int64(0), LValue(lua.LString("test")).ToInt64())
}

func TestLValueToInt(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, 0, LValue(nil).ToInt())
	assert.Equal(t, 0, LValue(lua.LNil).ToInt())

	// Test when the value is a number
	assert.Equal(t, 1, LValue(lua.LNumber(1)).ToInt())

	// Test when the value is a string (should return 0)
	assert.Equal(t, 0, LValue(lua.LString("test")).ToInt())
}

func TestLValueToUint8(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, uint8(0), LValue(nil).ToUint8())
	assert.Equal(t, uint8(0), LValue(lua.LNil).ToUint8())

	// Test when the value is a number
	assert.Equal(t, uint8(1), LValue(lua.LNumber(1)).ToUint8())

	// Test when the value is a string (should return 0)
	assert.Equal(t, uint8(0), LValue(lua.LString("test")).ToUint8())
}

func TestLValueToUint16(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, uint16(0), LValue(nil).ToUint16())
	assert.Equal(t, uint16(0), LValue(lua.LNil).ToUint16())

	// Test when the value is a number
	assert.Equal(t, uint16(1), LValue(lua.LNumber(1)).ToUint16())

	// Test when the value is a string (should return 0)
	assert.Equal(t, uint16(0), LValue(lua.LString("test")).ToUint16())
}

func TestLValueToUint32(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, uint32(0), LValue(nil).ToUint32())
	assert.Equal(t, uint32(0), LValue(lua.LNil).ToUint32())

	// Test when the value is a number
	assert.Equal(t, uint32(1), LValue(lua.LNumber(1)).ToUint32())

	// Test when the value is a string (should return 0)
	assert.Equal(t, uint32(0), LValue(lua.LString("test")).ToUint32())
}

func TestLValueToUint64(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, uint64(0), LValue(nil).ToUint64())
	assert.Equal(t, uint64(0), LValue(lua.LNil).ToUint64())

	// Test when the value is a number
	assert.Equal(t, uint64(1), LValue(lua.LNumber(1)).ToUint64())

	// Test when the value is a string (should return 0)
	assert.Equal(t, uint64(0), LValue(lua.LString("test")).ToUint64())
}

func TestLValueToBool(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, false, LValue(nil).ToBool())
	assert.Equal(t, false, LValue(lua.LNil).ToBool())

	// Test when the value is a boolean
	assert.Equal(t, true, LValue(lua.LBool(true)).ToBool())
	assert.Equal(t, false, LValue(lua.LBool(false)).ToBool())

	// Test when the value is a number (should return false)
	assert.Equal(t, false, LValue(lua.LNumber(1)).ToBool())
}

func TestLValueToFloat32(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, float32(0.0), LValue(nil).ToFloat32())
	assert.Equal(t, float32(0.0), LValue(lua.LNil).ToFloat32())

	// Test when the value is a float
	assert.Equal(t, float32(3.14), LValue(lua.LNumber(3.14)).ToFloat32())

	// Test when the value is an integer (should convert to float)
	assert.Equal(t, float32(1.0), LValue(lua.LNumber(1)).ToFloat32())
}

func TestLValueToFloat64(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, 0.0, LValue(nil).ToFloat64())
	assert.Equal(t, 0.0, LValue(lua.LNil).ToFloat64())

	// Test when the value is a float
	assert.Equal(t, 3.14, LValue(lua.LNumber(3.14)).ToFloat64())

	// Test when the value is an integer (should convert to float)
	assert.Equal(t, 1.0, LValue(lua.LNumber(1)).ToFloat64())
}

func TestLValueToTable(t *testing.T) {
	// Test when the value is nil
	assert.Equal(t, Table{&lua.LTable{}}, LValue(nil).ToTable())
	assert.Equal(t, Table{&lua.LTable{}}, LValue(lua.LNil).ToTable())

	// Test when the value is a table
	luaTable := &lua.LTable{}
	assert.Equal(t, LTable(luaTable), LValue(luaTable).ToTable())

	// Test when the value is a number (should return an empty table)
	assert.Equal(t, Table{&lua.LTable{}}, LValue(lua.LNumber(1)).ToTable())
}
