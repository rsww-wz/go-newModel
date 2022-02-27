package grammar

/*
带有类型参数的类型被叫做泛型类型

下面定义一个底层类型为切片的新类型,它是可以存储任何类型的切片
要使用泛型类型,要先对齐进行初始化,就是给类型指定一个实参
*/

type vector[T any] []T

func F2() {
	v1 := vector[int]{57, 178}
	PrintSlice(v1)

	v2 := vector[string]{"hello", "world"}
	PrintSlice(v2)
}
