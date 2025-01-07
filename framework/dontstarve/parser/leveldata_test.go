package dstparser

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevelDataOverridesCave(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/leveldataoverride.cave.lua")
	assert.Nil(t, err)
	overrides, err := ParseLevelDataOverrides(bytes)
	assert.Nil(t, err)

	assert.NotEmpty(t, overrides.Id)
	t.Log(overrides)
}

func TestParseLevelDataOverridesMaster(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/leveldataoverride.master.lua")
	assert.Nil(t, err)
	overrides, err := ParseLevelDataOverrides(bytes)
	assert.Nil(t, err)

	assert.NotEmpty(t, overrides.Id)
	t.Log(overrides)
}

func TestToMasterLevelDataOverridesLua(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/leveldataoverride.cave.lua")
	assert.Nil(t, err)
	overrides, err := ParseLevelDataOverrides(bytes)
	assert.Nil(t, err)

	assert.NotEmpty(t, overrides.Id)
	t.Log(overrides)

	overridesLua, err := ToMasterLevelDataOverridesLua(overrides)
	assert.Nil(t, err)
	t.Log(err)

	fmt.Println(string(overridesLua))
}

func TestToCaveLevelDataOverridesLua(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/leveldataoverride.cave.lua")
	assert.Nil(t, err)
	overrides, err := ParseLevelDataOverrides(bytes)
	assert.Nil(t, err)

	assert.NotEmpty(t, overrides.Id)
	t.Log(overrides)

	overridesLua, err := ToCaveLevelDataOverridesLua(overrides)
	assert.Nil(t, err)
	t.Log(err)

	fmt.Println(string(overridesLua))
}
