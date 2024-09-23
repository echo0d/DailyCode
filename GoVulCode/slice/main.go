package main

import "fmt"

func main() {
	a := make([]uint64, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	b := append(a, 4)
	// c := append(a, 5)

	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c)

}
