package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ComicMeta struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Img        string
	Title      string
	Day        string
}

func main() {
	// 省略了存json文件的步骤
	for _, queryWord := range os.Args[1:] {
		// 读取json文件
		workDir, _ := os.Getwd()
		outputFilePath := workDir + "/info.json"
		xkcdFile, err := os.Open(outputFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "opening error:%v", err)
		}
		// 给json解码成ComicMeta的数组
		var comicMetaSlice []ComicMeta
		if err = json.NewDecoder(xkcdFile).Decode(&comicMetaSlice); err != nil {
			fmt.Fprintf(os.Stderr, "decode error:%v", err)
		}
		// 按照Title找到对应的ComicMeta
		comicMeta := findByTitle(comicMetaSlice, queryWord)
		fmt.Printf("%s: %s\n", queryWord, comicMeta.Img)
		// 保存图片
		err = saveImageFromURL(comicMeta.Img, workDir+"/images/")
		if err != nil {
			return
		}
	}
}
func findByTitle(s []ComicMeta, title string) ComicMeta {
	for _, i := range s {
		// 找到了就返回ComicMeta数组的切片
		if title == i.Title {
			return i
		}
	}
	// 没找到返回一个空的ComicMeta
	return ComicMeta{}
}
func saveImageFromURL(url, distDir string) error {
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		return e
	}
	defer response.Body.Close()

	// open a file for writing
	// filepath.Base(url)是一个用于从URL中提取基本文件名（不包含路径）的函数。它是Go语言标准库path/filepath包中的一个函数。
	file, err := os.Create(distDir + filepath.Base(url))
	if err != nil {
		return e
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
