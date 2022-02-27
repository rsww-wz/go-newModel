package main

import "fmt"

/*
执行实例函数
	go test -run Example

可以直接在golang.org的godoc文档服务器上使用go playground运行实例代码

*/

func ExampleAdd() {
	fmt.Println(Add(3,6))
	fmt.Println(Add(20,34))
	// Output:
	// 9
	// 54
}

func ExampleSub() {
	fmt.Println(Sub(3,6))
	fmt.Println(Sub(100,34))
	// Output:
	// -3
	// 66
}

func ExampleMax() {
	fmt.Println(Max(3,6))
	fmt.Println(Max(20,34))
	// Output:
	// 6
	// 34
}

func ExampleMin() {
	fmt.Println(Min(3,6))
	fmt.Println(Min(20,34))
	// Output:
	// 3
	// 20
}



