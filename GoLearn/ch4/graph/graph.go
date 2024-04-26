package main

import "fmt"

// graph 这个 map 的key类型是一个字符串，value类型map[string]bool代表一个字符串集合
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	// NOTE: 此处在每个值首次作为key时才初始化
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	// graph[from][to]返回的是 from:map[to]的值，也就是addEdge中edges[to] = true。
	return graph[from][to]
}

func main() {
	addEdge("a", "a1")
	addEdge("a", "a2")
	addEdge("a", "a3")
	addEdge("b", "b1")
	addEdge("b", "b2")
	fmt.Printf("graph:\t%v\n", graph)
	fmt.Printf("hasEdge(\"a\", \"a3\"): %v\n", hasEdge("a", "a3"))
}
