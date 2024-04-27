package main

import "fmt"

func main() {
	//ages := make(map[string]int)
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	agesBob := ages["bob"]
	fmt.Println(agesBob) // 0
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages["bob"]) // 1
	fmt.Println(ages)        // map[alice:31 bob:1 charlie:34]
}
