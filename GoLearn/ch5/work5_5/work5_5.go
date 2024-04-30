package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		fmt.Println(words)
		fmt.Println(images)
	}

}

// CountWordsAndImages 函数执行一个HTTP GET请求获取HTML文档的URL，并返回其中的单词和图片数量。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages 函数递归遍历HTML的节点，统计words和images个数
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ws, is := countWordsAndImages(c)
		words += ws
		images += is
	}
	return words, images
}
