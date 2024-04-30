// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// 原本的for循环单独拿出来，改成一个递归函数
	return visitRecursion(links, n.FirstChild)

}

// visitRecursion 递归函数visitRecursion，用于遍历所有子节点
func visitRecursion(links []string, c *html.Node) []string {
	if c == nil {
		return links
	}
	links = visit(links, c)
	c = c.NextSibling
	return visitRecursion(links, c)
}
