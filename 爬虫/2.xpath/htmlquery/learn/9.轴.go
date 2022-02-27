package learn

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	req "learnhtmlquery/request"
)

func Axis() {
	url := "https://s.weibo.com/top/summary"
	doc := req.ParseHTML(url)
	divs, _ := htmlquery.QueryAll(doc, "//td/a")
	for _, item := range divs {
		//选取当前节点所有后代的文本
		lst, _ := htmlquery.QueryAll(item, "descendant::text()")
		for _, list := range lst {
			fmt.Println(list.Data)
		}
	}
}
