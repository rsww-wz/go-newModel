package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
)

func Index() {
	url := "https://www.baidu.com"
	doc := req.ParseHTML(url)

	// xpath的语法直接索引，下标是从1开始的
	item, _ := htmlquery.QueryAll(doc, "//div/a[1]")
	fmt.Println(htmlquery.OutputHTML(item[0], true))

	// 最后一个位置：[last()]
	item, _ = htmlquery.QueryAll(doc, "//div/a[last()]")
	fmt.Println(htmlquery.OutputHTML(item[0], true))
	fmt.Println()

	// 给定条件：[position()<K]
	// [position()<K]最后会转换成bool，如果是TRUE就截止，反之，下标还是从1开始
	// 6>5 -> true -> 1,2,3,4,5,6
	item, _ = htmlquery.QueryAll(doc, "//div/a[position() > 5]")
	for num, items := range item {
		content := htmlquery.OutputHTML(items, true)
		fmt.Println(num, content)
	}
}
