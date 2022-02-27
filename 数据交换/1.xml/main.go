package main

import (
	"fmt"
)

/*
command-line-arguments undefined
	因为是以文件的模式执行,所以会找不到定义
	即使是在命令行中，执行go run main.go也会报这样的错误

	命令行中执行,需要把依赖的文件全部在后面写上
	如果想简写,可以执行以下命令
		go run .
		go run ./
		go run *.go

	如果是IDE提供的功能，有以下方法
		配置以包的模式执行
		直接选中需要执行和依赖的文件，右键执行
		选中整个包，右键执行
*/

func main() {
	info := Person{
		Name:      "小明",
		Gender:    false,
		Age:       14,
		Telephone: "15119657172",
		Book: Book{
			Name:   "西游记",
			Author: "吴承恩",
			Price:  79.0,
		},
	}
	result := marshalStruct(&info)
	fmt.Println(result)
	fmt.Println()

	address := `E:\Document\0-code\Go\数据交换\1.xml\BookInfo.xml`
	//WriteStringXML(address, result)
	//WriteXML1(address,&info)

	var p Person
	unMarshalStruct(address, &p)
	fmt.Println(p)
	fmt.Println()

}
