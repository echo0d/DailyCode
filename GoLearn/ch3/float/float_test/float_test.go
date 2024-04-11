package float

import (
	"fmt"
)

// NOTE: 当整数大于23bit能表达的范围时，float32的表示将出现误差
func Example_float32() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	// Output:
	// true
}

func Example_floatvalue() {
	const e = 2.71828              // (approximately)
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数

	fmt.Println(e, Avogadro, Planck)
	// Output:
	// 2.71828 6.02214129e+23 6.62606957e-34
}
