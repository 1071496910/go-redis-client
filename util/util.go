package util

import "unsafe"

func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer((&s)))
}
