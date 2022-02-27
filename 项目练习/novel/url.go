package main

import (
	"github.com/antchfx/htmlquery"
	"log"
)

//func PageUrl() {
//	url := "https://ism399.com/story/p/1"
//}

func GetUrls(pageUrl string) []string{
	basePath := `https://ism399.com`
	urls := make([]string, 20)
	doc := Request(pageUrl)
	items, err := htmlquery.QueryAll(doc, `//div[@class="post-list"]/div/h2/a/@href`)
	if err != nil {
		log.Println(err)
	}

	for k, v := range items {
		url := htmlquery.OutputHTML(v, false)
		url = basePath + url
		urls[k] = url
	}

	return urls
}
