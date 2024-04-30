package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var linkMap = map[string]string{
	"a":      "href", // a标签href属性
	"img":    "src",  // img标签src属性
	"script": "src",  // script标签src属性
	"link":   "href", // link标签href属性
}

func main() {
	filePath := "golang.org.html"
	file, err := os.Open(filePath)
	if err != nil {
		// 处理打开文件时的错误
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	doc, err := html.Parse(file)
	//doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && linkMap[n.Data] != "" {
		for _, a := range n.Attr {
			if a.Key == linkMap[n.Data] {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
