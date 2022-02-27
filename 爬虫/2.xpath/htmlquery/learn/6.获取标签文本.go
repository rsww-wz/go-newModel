package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
)

func Text() {
	url := "https://www.baidu.com"
	doc := req.ParseHTML(url)

	// xpath语法 ：/text()，在要获取标签内容的后面加上/text()
	// 只能返回该标签下的直接文本，如果该标签下嵌套了其他标签并且有文本，不能获取
	items, _ := htmlquery.QueryAll(doc, "//div/a[@href]/text()")
	for _, item := range items {
		fmt.Print(item.Data)
	}
	fmt.Println("--------------------------------------------------")

	// 函数OutputHTML
	items, _ = htmlquery.QueryAll(doc, "//div/a[@href]")
	for _, item := range items {
		content := htmlquery.OutputHTML(item, false)
		fmt.Print(content)
	}
	fmt.Println("--------------------------------------------------")

	// 函数InnerText
	// 返回对象的开始和结束标记之间的文本，可以获取嵌套标签内的文本
	items, _ = htmlquery.QueryAll(doc, "//div/a[@href]")
	for _, item := range items {
		content := htmlquery.InnerText(item)
		fmt.Print(content)
	}


}
