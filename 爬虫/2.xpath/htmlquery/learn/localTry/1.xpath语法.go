package localTry

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"golang.org/x/net/html"
	"log"
	"runtime"
)

// 加载本地文档
func LoadFile() *html.Node {

	// 相对路径比较麻烦，用绝对路径
	var FilePath = `E:\Document\0-code\Go\爬虫\2.xpath\htmlquery\learn\localTry\test.txt`

	// runtime 表示调用该方法的路径，就是谁调用了runtime.Caller所在的方法的路径，返回的路径是唯一的
	fmt.Println(runtime.Caller(1))

	doc, err := htmlquery.LoadDoc(FilePath)
	if err != nil {
		log.Panicln("加载本地文档失败:", err)
	}
	return doc
}

// 获取标签
func GetTable(url ...string) {
	// 加载本地文件
	doc := LoadFile()

	// 加载网页
	//doc := req.ParseHTML(url[0])

	var expStr string

	for {
		fmt.Print("请输入:")
		_, _ = fmt.Scanln(&expStr)

		switch expStr {
		case "":
			break

		case "exit", "EXIT":
			return

		case "F5", "f5":
			if len(url) == 0 {
				GetTable()
			} else {
				GetTable(url[0])
			}

		case "in":
			var newUrl string
			for {
				fmt.Print("请输入新的网址:")
				_, _ = fmt.Scanln(&newUrl)
				if newUrl == "" {
					continue
				} else if newUrl == "exit"{
					return
				} else {
					GetTable([]string{newUrl}[0])
				}
			}

		default:
			_, err := xpath.Compile(expStr)

			// xpath编译失败
			if err != nil {
				log.Println("xpath路径编译失败:", err)
				break
			}

			// 能保证xpath是可以被编译的，所以忽略错误
			tables, _ := htmlquery.QueryAll(doc, expStr)

			//fmt.Println(len(tables))
			for k, v := range tables {
				item := htmlquery.OutputHTML(v, true)
				fmt.Printf("%d	%v\n", k+1, item)
			}
		}
	}
}
