package luax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func doScript(s string) (*VM, error) {
	vm := NewVM()
	err := vm.DoString(s)
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func TestTableGet(t *testing.T) {
	// Lua script with different types of values in a table
	script := `
	t = { 
		a = "test", 
		b = 42, 
		c = 3.14, 
		d = true, 
		e = { c=3.14 }, 
		f = 100, 
		g = 200.5, 
		h = false
	}`
	vm, err := doScript(script)
	assert.NoError(t, err)

	// Initialize the table from the global environment
	table := LTable(vm.G.Global).GetTable("t")

	// Test string retrieval
	getVal := table.Get("a")
	assert.NotNil(t, getVal)
	assert.Equal(t, "test", getVal.ToString())

	// Test integer retrieval as int8 (should be 100)
	getInt8 := table.GetInt8("f")
	assert.Equal(t, int8(100), getInt8)

	// Test integer retrieval as int16 (should be 100)
	getInt16 := table.GetInt16("f")
	assert.Equal(t, int16(100), getInt16)

	// Test integer retrieval as int32 (should be 100)
	getInt32 := table.GetInt32("f")
	assert.Equal(t, int32(100), getInt32)

	// Test integer retrieval as int (should be 100)
	getInt := table.GetInt("f")
	assert.Equal(t, 100, getInt)

	// Test integer retrieval as int64 (should be 100)
	getInt64 := table.GetInt64("f")
	assert.Equal(t, int64(100), getInt64)

	// Test unsigned integer retrieval as uint8 (should be 100)
	getUint8 := table.GetUint8("f")
	assert.Equal(t, uint8(100), getUint8)

	// Test unsigned integer retrieval as uint16 (should be 100)
	getUint16 := table.GetUint16("f")
	assert.Equal(t, uint16(100), getUint16)

	// Test unsigned integer retrieval as uint32 (should be 100)
	getUint32 := table.GetUint32("f")
	assert.Equal(t, uint32(100), getUint32)

	// Test unsigned integer retrieval as uint (should be 100)
	getUint := table.GetUint("f")
	assert.Equal(t, uint(100), getUint)

	// Test unsigned integer retrieval as uint64 (should be 100)
	getUint64 := table.GetUint64("f")
	assert.Equal(t, uint64(100), getUint64)

	// Test float retrieval as float32 (should be 200.5)
	getFloat32 := table.GetFloat32("g")
	assert.Equal(t, float32(200.5), getFloat32)

	// Test float retrieval as float64 (should be 200.5)
	getFloat64 := table.GetFloat64("g")
	assert.Equal(t, float64(200.5), getFloat64)

	// Test boolean retrieval as bool (should be false)
	getBoolFalse := table.GetBool("h")
	assert.Equal(t, false, getBoolFalse)

	// Test nested table retrieval
	getTable := table.GetTable("e")
	assert.NotNil(t, getTable)
	assert.True(t, len(getTable.ToMap()) > 0, "Expected not empty table")
}
