package grammar

import "fmt"

/*
[T any]参数的类型，意思是该函数支持任何T类型
在调用这个泛型函数的时候，可以显式指定类型参数
也可以省略

any是泛型的一个关键字
*/

// PrintSlice 打印任何类型的参数
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func F1() {
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	PrintSlice[float64]([]float64{1.1, 1.2, 1.6, 6.7})
	PrintSlice[string]([]string{"hello", "world", "good", "morning"})
	PrintSlice([]int64{55, 33, 22, 11})
}
