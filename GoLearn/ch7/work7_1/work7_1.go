// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*c++
	}
	return int(*c), sc.Err()
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*c++
	}
	return int(*c), sc.Err()
}

func main() {
	str := "hello world\nfoo bar\nbaz\n"
	var wc WordCounter
	wc.Write([]byte(str))
	fmt.Println("Word count:", wc)
	var lc LineCounter
	lc.Write([]byte(str))
	fmt.Println("Line count:", lc)

}
