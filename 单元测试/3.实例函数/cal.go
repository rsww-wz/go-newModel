package main

/*
实例函数
	以Example为前缀
	没有参数，也没有返回值

作用
	能够作为文档直接使用
	实例函数只要包含// output:
	也是可以通过测试的
*/


func Add(a,b int) int {
	return a + b
}

func Sub(a,b int) int {
	return a - b
}

func Max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a,b int) int {
	if a > b {
		return b
	}
	return a
}