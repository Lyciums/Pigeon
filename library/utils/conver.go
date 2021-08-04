package utils

import (
	"unsafe"
)

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}
