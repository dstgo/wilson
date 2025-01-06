package luax

import (
	"testing"

	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

func TestTableGet(t *testing.T) {
	script := `abc = { a = 1, b = false, c = "ka" }`
	state := lua.NewState()
	err := state.DoString(script)
	assert.Nil(t, err)

	val := LTable(state.G.Global).GetTable("abc")
	assert.NotNil(t, val)

	t.Log(val)
	getInt64 := val.GetInt64("a")
	getBool := val.GetBool("b")
	getString := val.GetString("c")

	assert.EqualValues(t, getInt64, 1)
	assert.EqualValues(t, getBool, false)
	assert.EqualValues(t, getString, "ka")
	t.Log(getInt64)
	t.Log(getBool)
	t.Log(getString)

	unknown := val.GetString("unknown")
	assert.EqualValues(t, unknown, "")
	t.Log(unknown)
}
