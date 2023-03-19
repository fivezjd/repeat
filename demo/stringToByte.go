package demo

import (
	"strings"
	"unsafe"
)

// 优化字符串转字节切片

func StringToByte(str *string) []byte {
	return *(*[]byte)(unsafe.Pointer(str))
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 1)“+”拼接适用于短小的、常量字符串（明确的，非变量），因为编译器会做优化。
// 4)builder拼接从性能和灵活性上都是上佳的选择。

// 优化builder

func builderString(p []string, cap int) string {
	var b strings.Builder
	// 减少扩容，所以可以提前计算好需要的容量
	b.Grow(cap)
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
