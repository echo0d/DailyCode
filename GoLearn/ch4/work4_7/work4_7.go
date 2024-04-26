//修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

package main

import (
	"fmt"
)

func main() {
	s := []byte("Hello 世界")

	fmt.Println(string(reverse(s)))
}

func reverse(bs []byte) []byte {
	// 先把byte数组转成rune数组
	runes := []rune(string(bs))
	// 然后正常些，和原本的reverse函数一样
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
