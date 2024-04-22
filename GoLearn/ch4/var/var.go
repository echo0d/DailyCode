package main

func main() {
	/*
		var a [3]int             // array of 3 integers
		fmt.Println(a[0])        // print the first element
		fmt.Println(a[len(a)-1]) // print the last element, a[2]

	*/
	/*
			// Print the indices and elements.
			for i, v := range a {
				fmt.Printf("%d %d\n", i, v)
			}

			// Print the elements only.
			for _, v := range a {
				fmt.Printf("%d\n", v)
			}

		var p [3]int = [3]int{1, 2, 3}
		var r [3]int = [3]int{1, 2}
		fmt.Println(p[1]) // "2"
		fmt.Println(r[2]) // "0"
	*/

	/*
		q := [...]int{1, 2, 3}
		fmt.Printf("%T\n", q) // "[3]int"

	*/
	/*	type Currency int

		const (
			USD Currency = iota // 美元
			EUR                 // 欧元
			GBP                 // 英镑
			RMB                 // 人民币
		)

		symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

		fmt.Println(RMB, symbol[RMB]) // "3 ￥"
	*/
	/*	r := [...]int{99: -1}
		fmt.Println(r)
	*/
	/*	a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}
		fmt.Println(a == b, a == c, b == c) // "true false false"
		//d := [3]int{1, 2}
		//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
	*/

}
