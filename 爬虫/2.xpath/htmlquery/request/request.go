package request

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
)

type UA struct {
	Key   string
	Value string
}

var UserAgent = UA{
	Key:   "User-Agent",
	Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.137 Safari/537.36",
}

// Get请求
// 用可变长参数实现可选参数
func Get(url string, ua ...UA) []byte {
	// 发起get请求
	resp, err := http.Get(url)
	defer func() { _ = resp.Body.Close() }()
	if err != nil {
		log.Panicln("请求失败", resp.StatusCode, err)
	}

	// 设置ua
	if len(ua) > 0 {
		resp.Header.Add(ua[0].Key, ua[0].Value)
	} else {
		resp.Header.Add(UserAgent.Key, UserAgent.Value)
	}

	// 获取网页内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("读取失败:", err)
	}

	//fmt.Println(string(content))

	return content
}

func ParseHTML(url string, ua ...UA) *html.Node {
	var content []byte
	if len(ua) > 0 {
		content = Get(url, ua[0])
	} else {
		content = Get(url)
	}

	doc, err := htmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		log.Panicln("解析失败:", err)
	}
	return doc
}

// test
func Request(url string) {
	req,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add(UserAgent.Key,UserAgent.Value)
	r,err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	content,_ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}
