package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func uintTest() {
	// uint32取值范围是0到4294967295
	var a uint32 = 2147483648
	var b uint32 = 2147483648
	var sum = a + b
	fmt.Println(reflect.TypeOf(sum))
	// uint32类型是无符号的，溢出时会循环回到0
	fmt.Printf("Sum is : %d\n", sum)
}

func intTest() {
	// int8取值范围-128到127
	var a int8 = 127
	var b int8 = 1
	var sum int8 = a + b
	fmt.Println(reflect.TypeOf(sum))
	// 溢出后的值会从-128开始循环，即128 % 256 = 128，所以sum的值会是-128
	fmt.Printf("Sum is : %d\n", sum)
}

func truncated() {
	var a int16 = 256
	var b = int8(a) // int8取值范围-128到127，256会溢出，取余后得到-128
	fmt.Println(b)
	//
	v, _ := strconv.Atoi("4294967377")
	s := int32(v)
	fmt.Println(s)
}

func main() {
	uintTest()
	intTest()
	truncated()
}
