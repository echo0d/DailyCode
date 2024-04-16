package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	// NOTE: bytes包还提供了Buffer类型用于字节slice的缓存。一个bytes.Buffer变量并不需要初始化，因为零值也是有效的
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
