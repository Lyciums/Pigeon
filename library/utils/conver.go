package utils

import (
	"strconv"
	"unsafe"
)

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
