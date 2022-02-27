package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
)

func Next() {
	url := "https://www.baidu.com"
	doc := req.ParseHTML(url)

	// 匹配所有div标签下面是a标签 最后匹配的是a标签
	divs, _ := htmlquery.QueryAll(doc, "//div/a")

	// 匹配标签的数量
	fmt.Println(len(divs))

	for _, items := range divs {
		itemHTML := htmlquery.OutputHTML(items, true)
		fmt.Println(itemHTML)
		fmt.Println()
	}
}
