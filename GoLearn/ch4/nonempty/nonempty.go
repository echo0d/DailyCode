package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	//fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data) // `["one" "three" "three"]`

}

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	// NOTE: strings[:0]表示一个原始数组strings的切片，但切片长度是0，容量是数组的cap
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
