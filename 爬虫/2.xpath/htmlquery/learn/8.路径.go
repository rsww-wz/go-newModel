package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
)

func Path() {
	url := "https://www.baidu.com"
	doc := req.ParseHTML(url)

	// | 表示或的意思，可以用在路径，属性等语法，所以可以表示多路径
	items, _ := htmlquery.QueryAll(doc, "//a/span | //a/img")
	for _, item := range items {
		content := htmlquery.OutputHTML(item, true)
		fmt.Println(content)
	}
	fmt.Println("----------------------------------------------")

	// .表示当前路径
	// ..表示当前节点的父节点，如果需要使用父节点，结构体自带有父节点属性
	items, _ = htmlquery.QueryAll(doc, "//div/a")
	for _, item := range items {
		// 用.表示当前节点，既//div/a
		// 在当前节点下搜索span[@class]
		list, _ := htmlquery.QueryAll(item, "./span[@class]")
		for _, lst := range list {
			content := htmlquery.OutputHTML(lst, true)
			fmt.Println(content)
		}
	}
}
