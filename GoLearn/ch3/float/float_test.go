package float

import "fmt"

// NOTE: 当整数大于23bit能表达的范围时，float32的表示将出现误差
func Example_float32() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	// Output:
	// true
}
