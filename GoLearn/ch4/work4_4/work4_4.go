package main

import (
	"fmt"
)

// 练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(s, 4))
}

func rotate(s []int, rotateTimes int) []int {
	//var result []int
	result := s
	for i := 0; i < rotateTimes; i++ {
		result = result[1:]
		result = append(result, result[0])
		//s = result
	}
	return result
}
