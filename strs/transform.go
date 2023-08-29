package strs

import "unsafe"

// Str2Byte 不安全
func Str2Byte(value string) []byte {
	b := unsafe.StringData(value)
	return unsafe.Slice(b, len(value))
}

// Str2Reverse 字符串翻转
func Str2Reverse(value string) string {
	return string(reverse([]rune(value)))
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
