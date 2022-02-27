package grammar

import "fmt"

/*
comparable约束的是所有可以比较的类型
comparable类型只能进行等于和不等于的比较

支持比较的类型
	布尔型
	数字类型
	字符串类型
	指针类型
	通道类型
	接口类型
	数组类型
	结构体类型


不支持比较的类型
	切片
	字典
*/

func Equal[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func F4() {
	arr1 := [3]int{1, 3, 5}
	arr2 := [3]int{1, 5, 3}
	fmt.Println(Equal[[3]int](arr1, arr2))
}
