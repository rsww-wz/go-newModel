package learn

import (
	"fmt"
	parse "github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
	"log"
)

func Root() {
	url := "http://www.baidu.com"
	doc := req.ParseHTML(url)

	// 匹配所有div标签
	divs, err := parse.QueryAll(doc, "//div")
	if err != nil {
		log.Panic("获取标签失败:", err)
	}

	// 获取所有div标签的数量
	fmt.Println(len(divs))

	// 每个div标签的内容
	// 第一个div标签是整个网页的内容，后面的div标签都是嵌套在第一个div标签的内容
	for _, items := range divs {
		itemHTML := parse.OutputHTML(items, true)
		fmt.Println(itemHTML)
		fmt.Println()
	}
}
