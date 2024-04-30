package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"regexp"
)

func main() {
	for _, url := range os.Args[1:] {
		doc, err := getDoc(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			os.Exit(1)
		}
		forEachNode(doc, startElement, endElement)
	}
}
func getDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %s", err)
	}
	return doc, nil
}

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attrs := ""
		for _, a := range n.Attr {
			attrs += fmt.Sprintf("%s=\"%s\" ", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s %s\\>\n", depth*2, "", n.Data, attrs)
		} else {
			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, attrs)
		}
		depth++
	}
	// html.CommentNode 注释节点
	if n.Type == html.CommentNode {
		fmt.Printf("%*s//%s\n", depth*2, "", n.Data)
	}
	// html.TextNode文本节点
	if n.Type == html.TextNode {
		// 删除字符串中空白字符
		nData := regexp.MustCompile(`[\n\s]+`).ReplaceAllString(n.Data, "")
		if nData != "" {
			fmt.Printf("%*s%s\n", depth*2, "", nData)
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
