package main

import "fmt"

func main() {
	//ages := make(map[string]int)
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	/*
		agesBob := ages["bob"]
		fmt.Println(agesBob) // 0
		ages["bob"] = ages["bob"] + 1
		fmt.Println(ages["bob"]) // 1
		fmt.Println(ages)        // map[alice:31 bob:1 charlie:34]
	*/
	/* 	// NOTE: 用range遍历map
	   	for name, age := range ages {
	   		fmt.Printf("%s\t%d\n", name, age)
	   	}
	   	// NOTE: 用sort对map排序
	   	var names []string
	   	// NOTE: 遍历map时候忽略value不用写_占位
	   	for name, _ := range ages {
	   		names = append(names, name)
	   	}
	   	sort.Strings(names)
	   	// NOTE: 遍历数组的时候index部分要用_占位
	   	for _, name := range names {
	   		fmt.Printf("%s\t%d\n", name, ages[name])
	   	}
	*/
	/*
		var ages map[string]int
		fmt.Println(ages == nil)    // "true"
		fmt.Println(len(ages) == 0) // "true"

	*/
	// NOTE: map的下标语法将产生两个值；第二个是一个布尔值，用于报告元素是否真的存在。
	age, ok := ages["bob"]
	if !ok {
		fmt.Printf("bob not found, age is %d\n", age)
	}
}
