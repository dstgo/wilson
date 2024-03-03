package strs

import "unsafe"

// Str2Byte converts a string to bytes slice directly without copying
func Str2Byte(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// Byte2Str converts a bytes slice to string directly without copying
func Byte2Str(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}
