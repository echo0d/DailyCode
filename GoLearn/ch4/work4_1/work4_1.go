package main

import (
	"fmt"
)

// NOTE: 练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。

type SHA256 [32]byte

func main() {
	md1, md2 := SHA256{14: 252, 31: 8}, SHA256{31: 5}
	bitDiff(&md1, &md2)
}

func bitDiff(md1, md2 *SHA256) int {
	diffCnt := 0
	fmt.Println("  sha1  \t  sha2 ")
	for i := range md1 {
		b1 := md1[i]
		b2 := md2[i]
		for i := 0; i < 8; i++ {
			// NOTE: get last bit of a byte
			lb1, lb2 := (b1>>i)&1, (b2>>i)&1
			if (lb1 ^ lb2) == 1 {
				diffCnt++
			}
		}
		fmt.Printf("%08b\t%08b\n", b1, b2)
	}
	fmt.Printf("bit diff count: %d\n", diffCnt)
	return diffCnt
}
