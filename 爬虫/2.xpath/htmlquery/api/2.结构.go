package api

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
)

func AttrsRun() {
	path := `E:\Document\0-code\Go\爬虫\2.xpath\htmlquery\api\test.txt`
	doc, err := htmlquery.LoadDoc(path)
	if err != nil {
		log.Fatalln("加载文档失败:", err)
	}

	divs,_ := htmlquery.QueryAll(doc,"//div")
	for _,v := range divs {
		content := htmlquery.OutputHTML(v,true)
		fmt.Println(content)
	}
	DocAttrs(divs[0])
	DocAttrs(divs[1])
}


func DocAttrs(node *html.Node) {
	fmt.Println("Type", node.Type)
	fmt.Println("Data", node.Data)
	fmt.Println("DataAtom", node.DataAtom)
	fmt.Println("Namespace", node.Namespace)
	fmt.Println("Attr", node.Attr)
	fmt.Println()
}
