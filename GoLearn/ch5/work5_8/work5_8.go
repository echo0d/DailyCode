package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	targetID := os.Args[2]
	if url == "" || targetID == "" {
		fmt.Fprintf(os.Stderr, "need params url and targetId")
	}
	doc, err := getDoc(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	foundNode := elementByID(doc, targetID)
	fmt.Printf("foundNode: %s%s\n", foundNode.Data, foundNode.Attr)
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

var foundNode *html.Node

func elementByID(doc *html.Node, id string) *html.Node {
	forEachNode(doc, id, findID, nil)
	return foundNode
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) {
	if pre != nil {
		goOn := pre(n, id)
		if !goOn {
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, id, pre, post)
	}
	if post != nil {
		post(n, id)
	}
}

func findID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				foundNode = n
				return false
			}
		}
	}
	return true
}
