package main

import (
	"fmt"
)

func OneTask(pageUrl string) {
	urls := GetUrls(pageUrl)
	for _, url := range urls {
		page := Request(url)
		content := GetContent(page)
		title := GetTile(page)
		WriteFile(content, title)
		fmt.Println(title, "爬取完成")
	}
}

func AllTask() {
	for i := 1; i <= 10; i++ {
		path := fmt.Sprintf("https://ism399.com/story/p/%d",i)
		OneTask(path)
	}
}
