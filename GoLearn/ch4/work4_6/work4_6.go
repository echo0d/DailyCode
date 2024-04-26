package main

import (
	"fmt"
	"unicode"
)

func main() {
	rs := []rune{'H', 'e', 'l', 'l', 'o', ' ', ' ', ' ', '世', '界'}
	fmt.Println("input string:\t", string(rs))
	bs := []byte(string(rs))
	fmt.Println("output string:\t", string(uniqueSpaceSlice(bs)))
}

func uniqueSpaceSlice(bs []byte) []byte {
	fmt.Println("input []bytes:\t", bs)
	for i := 0; i < len(bs); i++ {
		if unicode.IsSpace(rune(bs[i])) {
			bs = append(bs[:i], bs[i+1:]...)
			// 如果是空格就删掉，删掉以后就会长度变短，i要减1
			i--
		}
	}
	fmt.Println("output []bytes:\t", bs)
	return bs
}
