package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"
	fmt.Println(1i * 1i)             // "(-1+0i)", i^2 = -1
	xx := 1 + 2i
	yy := 3 + 4i
	fmt.Println(xx * yy)        // "(-5+10i)"
	fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"

}
