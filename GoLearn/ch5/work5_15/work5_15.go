package main

import (
	"fmt"
	"math"
)

func main() {
	maxNum, minNum, _ := MinMax(1, 2, 3)
	fmt.Print(minNum, maxNum)
	fmt.Println("\n")
	strsJoin, _ := Join("a", "bc", "ss")
	fmt.Print(strsJoin)
}

func MinMax(nums ...int) (max int, min int, err error) {
	if len(nums) == 0 {
		err = fmt.Errorf("err: 参数个数为0")
		return 0, 0, err
	}
	// 初始化最小值为最大整数
	min = math.MaxInt
	// 遍历所有参数，求最大值和最小值
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min, nil
}

func Join(strs ...string) (strsJoin string, err error) {
	if len(strs) == 0 {
		err = fmt.Errorf("err: 参数个数为0")
		return "", err
	}
	for _, str := range strs {
		strsJoin = strsJoin + str
	}
	return strsJoin, err
}
