package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

// NOTE: 练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
func main() {
	var shaType string
	// 命令参数shaType，默认SHA256
	flag.StringVar(&shaType, "shaType", "SHA256", "SHA256（default）or SHA384 or SHA512")
	flag.Parse()
	for _, s := range flag.Args() {
		resultStr := ""
		switch shaType {
		case "SHA256":
			c := sha256.Sum256([]byte(s))
			resultStr = fmt.Sprintf("%x", c)
		case "SHA384":
			c := sha512.Sum384([]byte(s))
			resultStr = fmt.Sprintf("%x", c)
		case "SHA512":
			c := sha512.Sum512([]byte(s))
			resultStr = fmt.Sprintf("%x", c)
		default:
			fmt.Printf("Hash Type %s 不支持，SHA256（default）or SHA384 or SHA512\n", shaType)
			os.Exit(1)
		}
		fmt.Printf("str:%s\tshaType:%s\t sha:%s\n", s, shaType, resultStr)
	}
}
