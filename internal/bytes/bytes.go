package bytes

import (
	"unsafe"
)

// StringToBytes BytesToString
// For more details, see https://github.com/golang/go/issues/53003#issuecomment-1140276077.

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
