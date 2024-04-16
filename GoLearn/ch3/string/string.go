package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	/*
		// NOTE: example1
		s := "hello, world"
		fmt.Println(len(s))     // "12"
		fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	*/
	/*
		// NOTE: error1
		c := s[len(s)]
		fmt.Println(c) // prints the null byte ('\u0000')
	*/

	/*
		// NOTE: example2
		fmt.Println(s[:5])             // "hello"
		fmt.Println(s[7:])             // "world"
		fmt.Println(s[:])              // "hello, world"
	*/
	/*
		// NOTE: example3
		fmt.Println("goodbye" + s[5:]) // "goodbye, world"
	*/
	/*
		// NOTE: example4
		ss := "left foot"
		t := ss
		ss += ", right foot"
		fmt.Println(ss) // "left foot, right foot"
		fmt.Println(t)  // "left foot"
	*/

	/*
	    // NOTE: error2
	   	// s[0] = 'L'      // compile error: cannot assign to s[0]
	   	// fmt.Println(s)
	*/
	/*
	   	// NOTE: example5
	    	const GoUsage = `Go is a tool for managing Go source code.

	   	Usage:
	   		go command [arguments]
	   	...`
	   	fmt.Println(GoUsage)
	*/
	/*
		// NOTE: example6
		world1 := "世界"
		world2 := "\xe4\xb8\x96\xe7\x95\x8c"
		world3 := "\u4e16\u754c"
		world4 := "\U00004e16\U0000754c"
		println(world1, world2, world3, world4)
	*/

	/*
		// NOTE: example7
		rune1 := '世'
		rune2 := '\u4e16'
		rune3 := '\U00004e16'
		println(rune1, rune2, rune3)
		fmt.Printf("%c,%c,%c\n", rune1, rune2, rune3)
	*/

	/*
		// NOTE: example8
		println(HasPrefix("hello", "hel"))
		println(HasSuffix("hello", "lo"))
		println(Contains("hello", "ll"))
	*/
	/*
		// NOTE: example9
		s := "美/ruːn/"
		countRunes(s)
	*/
	/*
		// NOTE: example10
		// "program" in Japanese katakana
		s := "プログラム"
		// （在Printf中的% x参数用于在每个十六进制数字前插入一个空格。）
		fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
		// 将[]rune类型转换应用到UTF8编码的字符串
		r := []rune(s)
		fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"
		fmt.Println(string(r)) // "プログラム"
		// 将一个整数转型为字符串
		fmt.Println(string(65))     // "A", not "65"
		fmt.Println(string(0x4eac)) // "京"
		// 如果对应码点的字符是无效的，则用\uFFFD无效字符作为替换
		fmt.Println(string(1234567)) // "?"
	*/
	// NOTE: example11
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	println(b, s2)
}
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func countRunes(s string) {
	fmt.Println(len(s)) // "13"
	// 方法1 直接调utf8.RuneCountInString()
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	// 方法2 range循环在处理字符串的时候，会自动隐式解码UTF8字符串
	n := 0

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		n++
	}
	println("number of runes:", n)
	// 方法3 遍历 调用utf8.DecodeRuneInString(s[i:])
	m := 0
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		i += size
		m++
	}
	println("number of runes:", m)
}
