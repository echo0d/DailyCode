package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// NOTE: 忽略value的map当作一个字符串集合
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		// 原本集合里不存在这个key就添加进去,value是bool类型
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
