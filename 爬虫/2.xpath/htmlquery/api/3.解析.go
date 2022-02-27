package api

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
)

func load() *html.Node {
	path := `E:\Document\0-code\Go\爬虫\2.xpath\htmlquery\api\test.txt`
	doc, err := htmlquery.LoadDoc(path)
	if err != nil {
		log.Fatalln("加载文档失败:", err)
	}

	return doc
}

func dealErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func myPrint(doc []*html.Node) {
	for _, v := range doc {
		content := htmlquery.OutputHTML(v, true)
		fmt.Println(content)
	}
}

func Parse(exp string) {
	doc := load()
	var n int
	for {
		fmt.Print("请输入:")
		_, _ = fmt.Scanln(&n)
		switch n {
		case 0:
			return

		case 1:
			// 返回匹配的第一个
			doc, err := htmlquery.Query(doc, exp)
			dealErr(err)
			fmt.Println(htmlquery.OutputHTML(doc, true))

		case 2:
			// 返回所有匹配所有的结果
			doc, err := htmlquery.QueryAll(doc, exp)
			dealErr(err)
			myPrint(doc)

		case 3:
			// 返回所有匹配所有的结果，但是如果表达式不能被解析，会报错
			doc := htmlquery.Find(doc, exp)
			myPrint(doc)

		case 4:
			//返回第一个匹配所有的结果，但是如果表达式不能被解析，会报错
			doc := htmlquery.FindOne(doc, exp)
			fmt.Println(htmlquery.OutputHTML(doc, true))

		default:
			fmt.Println("请重新输入:")
			break
		}
	}
}
