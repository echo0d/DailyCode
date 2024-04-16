package main

import "fmt"

func main() {
	s := "1234567890"
	fmt.Println(comma(s))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
