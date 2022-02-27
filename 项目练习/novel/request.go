package main

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func Request(url string) *html.Node {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln("构建失败:", err)
	}

	req.Header.Add("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.137 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("请求失败:", err)
	}

	defer func() { _ = resp.Body.Close() }()

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatalln("解析失败:", err)
	}
	return doc
}
