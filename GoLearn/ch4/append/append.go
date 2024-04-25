package main

import "fmt"

func main() {
	/* 	// rune类型 UTF-8 编码的 Unicode 码点，go语言中用它来处理中文
	   	var runes []rune
	   	for _, r := range "Hello, 世界" {
	   		runes = append(runes, r)
	   	}
	   	// Printf的%q把rune类型的Unicode码点转成单引号括起来的字符
	   	fmt.Printf("%q\n", runes)               // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
	   	fmt.Printf("%q\n", []rune("Hello, 世界")) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
	*/

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	/* 	// 内置appendx
	   	// NOTE: “...”省略号表示接收变长的参数为slice
	   	var x []int
	   	x = append(x, 1)
	   	x = append(x, 2, 3)
	   	x = append(x, 4, 5, 6)
	   	x = append(x, x...) // append the slice x
	   	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
	*/
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
