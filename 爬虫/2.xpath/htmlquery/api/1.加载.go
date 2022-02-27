package api

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"log"
)

func LoadFromFile() {
	path := `E:\Document\0-code\Go\爬虫\2.xpath\htmlquery\api\test.txt`
	doc, err := htmlquery.LoadDoc(path)
	if err != nil {
		log.Fatalln("加载文档失败:", err)
	}

	fmt.Printf("%T,%v\n", doc, doc)
}

// 内部是Get请求
func LoadFromUrl(url string) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Fatalln("从地址加载文档失败:", err)
	}
	fmt.Printf("%T,%v\n", doc, doc)
}

func LoadFromReader() {
	path := `E:\Document\0-code\Go\爬虫\2.xpath\htmlquery\api\test.txt`
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("文件读取失败:", err)
	}

	doc, err := htmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatalln("从接口加载文档失败:", err)
	}
	fmt.Printf("%T,%v\n", doc, doc)
}
