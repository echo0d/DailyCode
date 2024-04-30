package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	counts := map[string]int{}
	count(counts, doc)
	for tag, count := range counts {
		fmt.Printf("%s\t%d\n", tag, count)
	}
}

// outline 通过递归的方式遍历整个HTML结点树，并输出树的结构。
func count(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(counts, c)
	}
	return counts
}
