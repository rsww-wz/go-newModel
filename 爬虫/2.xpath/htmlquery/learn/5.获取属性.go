package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
)

func Attrs() {
	url := "https://www.baidu.com"
	doc := req.ParseHTML(url)
	// 获取所有a标签的href属性
	items, _ := htmlquery.QueryAll(doc, "//a/@href")

	// 使用OutputHTML函数,直接输出获取到的结果，需要在xpath语法中指定属性
	// xpath语法可以匹配多个属性
	for _, item := range items {
		attrs := htmlquery.OutputHTML(item, false)
		fmt.Println(attrs)
	}
	fmt.Println("-----------------------------------------------------------")
	fmt.Println()

	// 使用SelectAttr函数，xpath语法只需要确定路径，属性在函数中指定
	// SelectAttr函数无法匹配多个属性
	items, _ = htmlquery.QueryAll(doc, "//span")
	for _, item := range items {
		attrs := htmlquery.SelectAttr(item, "class")
		fmt.Println(attrs)
	}
}
