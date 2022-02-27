package project

import (
	"fmt"
	"github.com/antchfx/htmlquery"
)

func Sina() {
	url := `https://s.weibo.com/top/summary`
	doc := request(url)
	tables, _ := htmlquery.QueryAll(doc, `//div`)
	for k, v := range tables {
		content := htmlquery.OutputHTML(v, false)
		fmt.Println(k, content)
	}
}
