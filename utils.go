package stream

import (
	"unsafe"
)

// If string is readonly, modifying bytes will cause panic
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// The change of bytes will cause the change of string synchronously,
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// BitToU8(0b10010110, 0, 3) == 0b00000110
func BitToU8(b byte, startBit int8, endBit int8) uint8 {
	b <<= 7 - endBit
	b >>= 7 - endBit + startBit
	return b
}
