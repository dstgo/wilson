package dstparser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseModInfoLua(t *testing.T) {
	dir := "testdata/workshop"
	entries, err := os.ReadDir(dir)
	assert.Nil(t, err)

	for _, entry := range entries {
		t.Log(entry.Name())
		infopath := filepath.Join(dir, entry.Name(), "modinfo.lua")
		bytes, err := os.ReadFile(infopath)
		if os.IsNotExist(err) {
			continue
		}
		assert.Nil(t, err)

		modInfo, err := ParseModInfoWithEnv(bytes, fmt.Sprintf("workshop-%s", entry.Name()), "zh")
		t.Log(modInfo.Name, modInfo.Author, modInfo.Version)
	}
}

func TestParseModOverridesLua(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/modoverrides.lua")
	assert.Nil(t, err)

	overrides, err := ParseModOverrides(bytes)
	assert.Nil(t, err)

	t.Log(len(overrides))
}

func TestFromModOverrideOption(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/modoverrides.lua")
	assert.Nil(t, err)

	overrides, err := ParseModOverrides(bytes)
	assert.Nil(t, err)

	overrideLua, err := ToModOverrideLua(overrides)
	assert.Nil(t, err)

	fmt.Println(overrideLua)
}
