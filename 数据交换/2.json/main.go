package main

import (
	"encoding/json"
	"fmt"
	myJson "json/jsoniter"
	parse "json/jsonparser"
)

type Book struct {
	Name   string
	Author string
	Price  float64
}

type Person struct {
	Name      string
	Gender    bool
	Age       int
	Telephone string
	Book      Book
}

func main() {
	info := Person{
		Name:      "小红",
		Gender:    true,
		Age:       20,
		Telephone: "13638562693",
		Book: Book{
			Name:   "红楼梦",
			Author: "曹雪芹",
			Price:  85.6,
		},
	}

	str1 := marsh(&info)
	str2 := formatMarsh(&info)
	fmt.Println(string(str1))
	fmt.Println(string(str2))
	fmt.Println()

	var p1 Person
	var p2 Person
	unMarsh(&p1,str1)
	unMarsh(&p2,str2)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println()

	fmt.Println("第三方库")
	myJson.JsoniterMarsh(&info)
	fmt.Println()

	var p3 Person
	var p4 Person
	var p5 Person
	myJson.JsoniterUnMarsh(&p3,str1)
	myJson.JsoniterUnMarsh(&p4,str2)
	fmt.Println(p3)
	fmt.Println(p4)

	myJson.UnMarshJsonStr(&p5,string(str2))
	fmt.Println(p5)
	fmt.Println()

	myJson.Test()
	fmt.Println()

	fmt.Println("-------------jsonparser------------------")
	parse.GetType()
	parse.GetValue()
	parse.Modify()
	parse.Remove()
	fmt.Println()
	parse.ParseVal()
}

// 序列化
func marsh(v interface{}) []byte {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return jsonStr
}

// 格式化序列化
func formatMarsh(v interface{}) []byte {
	formatStr, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}

	return formatStr
}

// 不管json有没有格式化，都可以转换成对应的结构体
func unMarsh(v interface{}, data []byte) {
	err := json.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
}
