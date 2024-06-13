// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
package main

import (
	"bufio"
	"bytes"
	"io"
)

type countWriter struct {
	w io.Writer
	c *int64
}

func (c countWriter) Write(p []byte) (n int, err error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		n++
		*c.c++
	}
	return int(n), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	return &countWriter{w, &c}, &c
}

func main() {

}
