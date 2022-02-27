package main

import (
	"fmt"
	"os"
)

func main() {
	args()
}

// 函数获取命令行参数
func args() {
	fmt.Println("命令行的参数有:", len(os.Args))

	// 遍历os.Args切片，就可以得到所以的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

	// 字符串切片
	if len(os.Args) > 1 {
		fmt.Printf("args = %v,%T\n", os.Args, os.Args[1])
	}
}
