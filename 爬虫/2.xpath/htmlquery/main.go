package main

import (
	"fmt"
	a "learnhtmlquery/api"
	practise "learnhtmlquery/learn/localTry"
	p "learnhtmlquery/project"
	req "learnhtmlquery/request"
)

func main() {
	//learn_api()
	//learn_xpath()
	//test()
	//learn.Root()
	//learn.Root()
	project()
}

func project() {
	//p.Sina()
	p.GetTotalInfo()

}

func test() {
	//url := "https://studygolang.com/articles/21918"
	//url := "https://s.weibo.com/top/summary"
	url := "https://baidu.com"
	content := req.Get(url, req.UserAgent)
	fmt.Println(string(content))
	//err := ioutil.WriteFile("./test.txt", content, 0777)
	//if err != nil {
	//	log.Panicln("文件写入失败", err)
	//}
}

func learn_xpath() {
	//url := `https://movie.douban.com/top250`
	//url := `http://www.baidu.com`
	//practise.GetTable(url)
	practise.GetTable()
}

func learn_api() {
	//url := `http://www.baidu.com`
	//a.LoadFromFile()
	//a.LoadFromReader()
	//a.LoadFromUrl(url)

	//a.AttrsRun()
	//a.Parse("//div")
	a.GetContent("//a")
}
