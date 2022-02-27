package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	req "learnhtmlquery/request"
)

func AllAttrs(doc *html.Node) {
	url := "https://www.baidu.com"
	doc = req.ParseHTML(url)

	// 匹配所有父标签是//div/a，子标签是拥有任何属性的span标签，如果span标签没有属性，不会匹配
	divs, _ := htmlquery.QueryAll(doc, "//div/a/span[@*]")
	fmt.Println(len(divs))

	for _, items := range divs {
		itemHTML := htmlquery.OutputHTML(items, true)
		fmt.Println(itemHTML)
		fmt.Println()
	}
}

func AllTables(doc *html.Node) {
	url := "https://www.baidu.com"
	doc = req.ParseHTML(url)

	// 语法：*表示匹配任何元素，无论是标签还是通配符
	// 匹配父标签是div/a这样结构的标签下面的所有标签

	divs, _ := htmlquery.QueryAll(doc, "//div/a/*")
	fmt.Println(len(divs))

	for _, items := range divs {
		itemHTML := htmlquery.OutputHTML(items, true)
		fmt.Println(itemHTML)
		fmt.Println()
	}
}
