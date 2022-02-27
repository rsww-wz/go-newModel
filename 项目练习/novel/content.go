package main

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
	"strings"
)

func GetContent(page *html.Node) string{
	item,err := htmlquery.Query(page,`//div[@class="topic-content"]/p`)
	if err != nil {
		log.Println(err)
	}

	content := htmlquery.OutputHTML(item,false)
	content = strings.Replace(content,"<br/>","\n",-1)
	content = strings.Replace(content,"ism301.com","",-1)
	content = strings.Replace(content,"聽","",-1)
	return content
}

func GetTile(page *html.Node) string{
	item,err := htmlquery.Query(page,`//div/h1`)
	if err != nil {
		log.Println(err)
	}
	title := htmlquery.OutputHTML(item,false)
	return title
}