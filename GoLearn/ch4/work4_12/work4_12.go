package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	filePath := "info.json"
	checkAndCreateJSONFile(filePath)
	for _, queryWord := range os.Args[1:] {
		workDir, _ := os.Getwd()
		outputFilePath := workDir + "/info.json"
		xkcdFile, err := os.Open(outputFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "opening error:%v", err)
		}

		var comicMetaSlice []ComicMeta
		if err = json.NewDecoder(xkcdFile).Decode(&comicMetaSlice); err != nil {
			fmt.Fprintf(os.Stderr, "decode error:%v", err)
		}
		comicMeta := findByTitle(comicMetaSlice, queryWord)
		fmt.Printf("%s: %s\n", queryWord, comicMeta.Img)
	}
}

func checkAndCreateJSONFile(filePath string) {
	// 判断文件是否存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 文件不存在，创建文件
		err := createJSONFile(filePath)
		if err != nil {
			fmt.Printf("Failed to create JSON file: %v\n", err)
			return
		}
		fmt.Println("JSON file created successfully.")
	} else if err != nil {
		fmt.Printf("Failed to check file existence: %v\n", err)
		return
	}

	// 文件存在，继续其他操作
	fmt.Println("JSON file exists.")
}

func createJSONFile(filePath string) error {
	var urls []string
	for i := 1; i <= 100; i++ {
		urls = append(urls, fmt.Sprintf("https://com/%d/info.0.json", i))
	}

	// 创建一个空的ComicMeta切片
	var ComicMetas []ComicMeta

	// 循环遍历URL列表
	for _, url := range urls {
		// 发送HTTP GET请求获取JSON数据
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to get JSON data from %s: %v\n", url, err)
			continue
		}
		defer response.Body.Close()

		// 读取响应的JSON数据
		jsonData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Failed to read JSON data from %s: %v\n", url, err)
			continue
		}

		// 解析JSON数据并将其存储在ComicMeta结构体中
		var ComicMeta ComicMeta
		err = json.Unmarshal(jsonData, &ComicMeta)
		if err != nil {
			fmt.Printf("Failed to parse JSON data from %s: %v\n", url, err)
			continue
		}

		// 将ComicMeta结构体添加到切片中
		ComicMetas = append(ComicMetas, ComicMeta)
	}

	// 将切片中的JSON数据转换为JSON字符串
	jsonData, err := json.MarshalIndent(ComicMetas, "", "  ")
	if err != nil {
		fmt.Printf("Failed to convert data to JSON: %v\n", err)
		return err
	}

	// 将JSON字符串写入文件
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Failed to write JSON file: %v\n", err)
		return err
	}

	fmt.Println("JSON data added to file successfully.")
	return err
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
