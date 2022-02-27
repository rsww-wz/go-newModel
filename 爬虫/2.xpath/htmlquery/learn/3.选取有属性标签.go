package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
)

func TableWithAttr() {
	url := "https://www.baidu.com"
	doc := req.ParseHTML(url)

	// 匹配一个属性并指定属性值
	// 语法：[@属性 = \"属性\"] 最外层可以使用反引号

	divs, _ := htmlquery.QueryAll(doc, "//a[@target=\"_blank\"]")
	for _, items := range divs {
		itemHTML := htmlquery.OutputHTML(items, false)
		fmt.Println(itemHTML)
		fmt.Println("-----------------------------")
	}

	// 匹配多个属性
	// 语法：`[@属性1 and @属性2]`
	divs, _ = htmlquery.QueryAll(doc, "//a[@id and @class]")
	for _, items := range divs {
		item := htmlquery.OutputHTML(items, true)
		fmt.Println(item)
	}
	fmt.Println("---------------------------------------------------")

	// 匹配多个属性中的一个
	// 语法 : `[@属性1 | @属性2]`
	divs, _ = htmlquery.QueryAll(doc, "//span[@id | @class]")
	for _, items := range divs {
		item := htmlquery.OutputHTML(items, true)
		fmt.Println(item)
	}
}
