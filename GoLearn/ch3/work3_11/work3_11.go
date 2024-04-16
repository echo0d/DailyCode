package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	testString := "-152215.23223456"
	fmt.Println(add(testString))
}

func add(s string) string {
	// 处理符号
	symbol := s[0]
	s = s[1:]
	// 处理浮点数
	dotIndex := strings.LastIndex(s, ".")
	decimals := s[dotIndex:] // 小数部分
	println(decimals)
	s = s[:dotIndex]
	println(s)
	// 整数部分buffer
	var buf1 bytes.Buffer
	buf1.WriteByte(symbol)
	a := comma(buf1, s)
	println(a)
	// 小数部分buffer
	var buf2 bytes.Buffer
	b := comma(buf2, decimals)
	println(b)
	return (a + b)
}

func comma(buf bytes.Buffer, s string) string {
	runeCnt := utf8.RuneCountInString(s)
	sep := runeCnt % 3
	cnt := 0
	for _, r := range s {
		// r is rune
		if cnt%3 == sep && cnt != 0 { // 每个逗号的位置为 sep + 3n
			buf.WriteString(",")
		}
		buf.WriteRune(r)
		cnt++
	}
	return buf.String()
}
