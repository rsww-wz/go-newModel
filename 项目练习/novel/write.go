package main

import (
	"io/ioutil"
	"log"
)

func WriteFile(content,title string) {
	path := `E:\多媒体\文档\爬虫获取\`
	path = path + title + ".txt"

	err := ioutil.WriteFile(path,[]byte(content),0777)
	if err != nil {
		log.Println("文件写入失败",err)
	}
}
