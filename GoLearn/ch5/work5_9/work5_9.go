package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, input := range os.Args[1:] {
		fmt.Println("expand: ", expand(input, allAdd1))
	}
}

func expand(s string, f func(string) string) string {
	// strings.Replace() ，下面的n表示替换前n个，n<0表示替换数量不限制
	ret := strings.Replace(s, "foo", f("foo"), -1)
	return ret
}
func add1(r rune) rune {
	return r + 1
}

func allAdd1(s string) string {
	return strings.Map(add1, s)
}
