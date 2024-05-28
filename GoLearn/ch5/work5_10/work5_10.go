package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i, course)
	}
}

func topoSort(m map[string][]string) map[int]string {
	//var order []string
	var order = make(map[int]string)
	seen := make(map[string]bool)
	// NOTE: 当匿名函数需要被递归调用时，我们必须首先声明一个变量visitAll 再将匿名函数赋值给这个变量
	var visitAll func(items []string)
	index := 1
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[index] = item
				index++
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	visitAll(keys)
	return order
}
