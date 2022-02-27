package project

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Movies struct {
	Rank      int
	CnName    string
	EnName    string
	Country   string
	Class     string
	Year      string
	Score     float64
	Directors string
	Actors    string
}

func NewMovies(n int, Country, Class, Year, Directors, Actors string) *Movies {
	return &Movies{
		Rank:      n,
		Country:   Country,
		Class:     Class,
		Year:      Year,
		Directors: Directors,
		Actors:    Actors,
	}
}

func request(url string) *html.Node {
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

func GetTotalInfo() {
	url := "https://movie.douban.com/top250"
	doc := request(url)
	cn, en := GetTitle(doc)
	infos := GetInfo(doc)
	data := ClearInfo(infos)
	combineInfo := CombineInfo(data)
	combineInfo = CombineInfo(data)
	combineInfo = TotalCombine(cn, en, combineInfo)
	l := len(combineInfo)
	for i := 0; i < l; i++ {
		PrintData(combineInfo[i])
	}
	//test()
}

func GetTitle(doc *html.Node) ([]string, []string) {
	var cn = make([]string, 0, 50)
	var en = make([]string, 0, 50)

	tables, _ := htmlquery.QueryAll(doc, `//div[@class="hd"]/a[@href]`)
	for _, v := range tables {
		cnTitle, _ := htmlquery.Query(v, "./span[1]/text()")
		enTitle, _ := htmlquery.Query(v, "./span[2]/text()")
		cnName := htmlquery.OutputHTML(cnTitle, true)
		enName := htmlquery.OutputHTML(enTitle, true)
		cn = append(cn, cnName)
		en = append(en, enName)
	}
	return cn, en
}

func GetInfo(doc *html.Node) []string {
	var info = make([]string, 0, 50)

	tables, _ := htmlquery.QueryAll(doc, `//div[@class="info"]/div[@class="bd"]`)
	for _, v := range tables {
		infos, _ := htmlquery.Query(v, "./p[1]")
		content := htmlquery.OutputHTML(infos, false)
		info = append(info, content)
	}

	return info
}

func ClearInfo(infos []string) map[string][]string {
	var (
		year      = make([]string, 0, 50)
		country   = make([]string, 0, 50)
		class     = make([]string, 0, 50)
		actors    = make([]string, 0, 50)
		directors = make([]string, 0, 50)
		info      = make(map[string][]string)
	)

	for _, v := range infos {
		rYear := regexp.MustCompile(`[\d]+`)
		rDirector := regexp.MustCompile(`导演: (.*?)主演:`)
		rActor := regexp.MustCompile(`导演.*?主演: (.*)`)
		rCountry := regexp.MustCompile(`(?s:.*?[\d]+.*/(.*)/)`)
		rClass := regexp.MustCompile(`(?s:.*?[\d]+.*/.*/(.*)\n)`)

		r1 := rYear.FindAllString(v, -1)
		r2 := rDirector.FindAllStringSubmatch(v, -1)
		r3 := rActor.FindAllStringSubmatch(v, -1)
		r4 := rCountry.FindAllStringSubmatch(v, -1)
		r5 := rClass.FindAllStringSubmatch(v, -1)

		year = append(year, r1[0])
		directors = append(directors, r2[0][1])
		actors = append(actors, strings.Replace(r3[0][1], "...<br/>", "", -1))
		country = append(country, strings.Replace(r4[0][1], " ", "", -1))
		class = append(class, strings.Trim(r5[0][1], " "))
	}

	info["year"] = year
	info["directors"] = directors
	info["actors"] = actors
	info["country"] = country
	info["class"] = class

	return info
}

func CombineInfo(data map[string][]string) []*Movies {
	length := len(data["year"])
	var dataLst = make([]*Movies, 0, 50)
	for i := 0; i < length; i++ {
		info := NewMovies(i+1, data["country"][i], data["class"][i], data["year"][i], data["directors"][i], data["actors"][i])
		dataLst = append(dataLst, info)
	}
	return dataLst
}

func TotalCombine(cn, en []string, data []*Movies) []*Movies {
	length := len(data)
	for i := 0; i < length; i++ {
		data[i].CnName = cn[i]
		data[i].EnName = en[i]
	}

	return data
}

func PrintData(data *Movies) {
	fmt.Printf("%d\t%s\n%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t\n",
		data.Rank, data.CnName, data.EnName, data.Year, data.Class, data.Country, data.Actors, data.Directors)
}

func test() {
	str := `
                            导演: 弗兰克·德拉邦特 Frank Darabont   主演: 蒂姆·罗宾斯 Tim Robbins /...<br/>
                            1994 / 美国 / 犯罪 剧情
`

	re := regexp.MustCompile(`(?s:导演: (.*)主演: (.*?)/.*<br/>.* ([\d]+) / ([\W]+) / (.*))`)
	res := re.FindAllStringSubmatch(str, -1)
	for k, v := range res[0] {
		fmt.Println(k, v)
	}
}
