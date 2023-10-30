package vax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {

	rule := RangeLength(1, 10, false)
	runeRule := RangeLength(1, 10, true)

	str1 := "01234567899"
	str2 := "1234567"
	str3 := "你好你好你好你好你好你好你好"
	str4 := "你好你好你好"

	// #1
	{
		assert.NotNil(t, rule.Validate("", str1))
		assert.Nil(t, rule.Validate("", str2))
		assert.NotNil(t, rule.Validate("", str3))
		assert.NotNil(t, rule.Validate("", str4))
	}

	// #2
	{
		assert.NotNil(t, runeRule.Validate("", str1))
		assert.Nil(t, runeRule.Validate("", str2))
		assert.NotNil(t, runeRule.Validate("", str3))
		assert.Nil(t, runeRule.Validate("", str4))
	}
}
