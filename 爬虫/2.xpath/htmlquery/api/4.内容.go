package api

import (
	"fmt"
	"github.com/antchfx/htmlquery"
)

func GetContent(exp string) {
	doc, err := htmlquery.Query(load(), exp)
	dealErr(err)

	var n int
	for {
		fmt.Print("请输入:")
		_, _ = fmt.Scanln(&n)
		switch n {
		case 0:
			return

		case 1:
			content := htmlquery.OutputHTML(doc, true)
			fmt.Println(content)

		case 2:
			content := htmlquery.OutputHTML(doc, false)
			fmt.Println(content)

		case 3:
			content := htmlquery.InnerText(doc)
			fmt.Println(content)

		case 4:
			var addr string
			fmt.Print("请输入要获取的属性:")
			_, _ = fmt.Scanln(&addr)
			content := htmlquery.SelectAttr(doc, addr)
			fmt.Println(content)

		default:
			fmt.Println("请重新输入:")
			break
		}
	}
}
