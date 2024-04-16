package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(containCharsEqual("hahaqwe", "ahewqah"))
}

func containCharsEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// NOTE: 统计字符出现次数，放到两个map里
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}
	// NOTE: 新函数 判断是否深度相等
	return reflect.DeepEqual(m1, m2)
}
