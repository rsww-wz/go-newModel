package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义参数方法1，返回指针
	var user = flag.String("user","root","用户名")
	fmt.Println(*user)

	// 定义参数方法2，绑定变量 (推荐)
	var pwd string
	flag.StringVar(&pwd,"pwd","","密码")
	fmt.Println(pwd)


	// 解析参数
	flag.Parse()

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())

	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())

	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}