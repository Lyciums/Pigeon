package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum(StringToBytes(s)))
}

// Base64Encode 编码
func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Base64EncodeByte 编码成字节数组
func Base64EncodeByte(b []byte) string {
	return Base64Encode(b)
}

// Base64EncodeString 编码成字符串
func Base64EncodeString(s string) string {
	return Base64Encode([]byte(s))
}

// Base64Decode 解码
func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Base64DecodeToString 解码成字符串
func Base64DecodeToString(s string) string {
	if v, err := Base64Decode(s); err == nil {
		return string(v)
	}
	return ""
}

// Base64DecodeToByte 解码为字节数组
func Base64DecodeToByte(s string) []byte {
	return []byte(Base64DecodeToString(s))
}

// Base64DecodeByteToByte 解码字节数组到字节数组
func Base64DecodeByteToByte(b []byte) []byte {
	return Base64DecodeToByte(string(b))
}

// Base64DecodeStringToByte 解码字符串到字节数组
func Base64DecodeStringToByte(s string) []byte {
	return Base64DecodeToByte(s)
}

func IsBase64(s string) bool {
	b, _ := Base64Decode(s)
	return Base64EncodeString(string(b)) == s
}

func IsBase64String(s string) bool {
	return IsBase64(s)
}

func IsBase64Byte(b []byte) bool {
	return IsBase64(string(b))
}
