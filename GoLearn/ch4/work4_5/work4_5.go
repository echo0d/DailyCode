package main

import (
	"fmt"
)

// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func main() {
	s := []string{"c", "a", "a", "a", "i"}
	fmt.Println(uniqueSlice(s))
}

func uniqueSlice(strSlice []string) []string {
	tempStr := ""
	for i := 0; i < len(strSlice); i++ {
		if tempStr == strSlice[i] {
			strSlice = append(strSlice[:i], strSlice[i+1:]...)
			// 重复的话，长度要减一了
			i--
		}
		tempStr = strSlice[i]
	}
	return strSlice
}
