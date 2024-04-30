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
	outline(nil, doc)
}

// outline 通过递归的方式遍历整个HTML结点树，并输出树的结构。
// 当outline调用自身时，被调用者接收的是stack的拷贝。被调用者对stack的元素追加操作，这个过程并不会修改调用方的stack。
// 因此当函数返回时，调用方的stack与其调用自身之前完全一致。
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
