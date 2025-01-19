package strs

import (
	"unsafe"
)

func BytesToString(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
