package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "asd..zqwoirn1世界11099"
	fmt.Println(comma(s))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	// NOTE: 这里直接写len不行 碰到中文就不对了
	runeCnt := utf8.RuneCountInString(s)
	sep := runeCnt % 3
	// NOTE: 需要定义一个cnt，然后进循环里++，要直接用range s 的index结果不对,因为每次遇到中文index加2
	cnt := 0
	for _, r := range s {
		// r is rune
		println(cnt, string(r))
		if cnt%3 == sep && cnt != 0 { // 逗号的位置为 sep + 3n
			buf.WriteString(",")
		}
		buf.WriteRune(r)
		cnt++
	}
	return buf.String()
}
