package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	strings := []string{
		"/a/./../b",
		"a/../b",
		"..a/./../b",
		"a/../../../../../../../b/c",
	}

	domain := "ytm.com"
	fmt.Println("结果1如下：")
	for _, s := range strings {
		fmt.Println(filepath.Join(domain, s))
	}
	fmt.Println("结果2如下：")
	for _, s := range strings {
		fmt.Println(filepath.Clean(filepath.Join(domain, s)))
	}
}
