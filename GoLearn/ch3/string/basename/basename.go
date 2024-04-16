// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename1("a"))
	fmt.Println(basename1("a.go"))
	fmt.Println(basename1("a/b/c.go"))
	fmt.Println(basename2("a"))
	fmt.Println(basename2("a.go"))
	fmt.Println(basename2("a/b/c.go"))
}

// NOTE: 第一个版本并没有使用任何库，全部手工硬编码实现
func basename1(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// NOTE: 这个简化版本使用了strings.LastIndex库函数
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
