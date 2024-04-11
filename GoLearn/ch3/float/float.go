package main

import (
	"fmt"
	"math"
)

func main() {
	// PrintFloat()
	PrintNAN()

}

func PrintFloat() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func PrintNAN() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

}
