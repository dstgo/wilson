package strs

import (
	"unsafe"
)

// BytesToString convert []byte to string without memory allocate,
// it should be careful to use.
func BytesToString(bs []byte) string {
	// G103: Use of unsafe calls should be audited
	// #nosec
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}

// StringToBytes convert string to []byte without memory allocate,
// it should be careful to use.
func StringToBytes(s string) []byte {
	// G103: Use of unsafe calls should be audited
	// #nosec
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
