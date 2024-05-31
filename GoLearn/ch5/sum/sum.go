package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	fmt.Printf("%T\n", f)       // "func(...int)"
	fmt.Printf("%T\n", g)       // "func([]int)"
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
}
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func f(...int) {}
func g([]int)  {}
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
