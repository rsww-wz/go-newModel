package grammar

import (
	"constraints"
	"fmt"
)

/*
约束类型
	约束哪些类型是可以被用做泛型的

如何定义约束类型
	约束类型是一个接口
	新的运算符 ~
	用于约束类型

官方已经定义好的约束类型
在constraints包里面
	Signed 	  有符号类型
	Unsigned  无符号类型
	Integer   整数类型
	Float	  浮点类型
	Complex	  复数类型
	Ordered	  有序类型 (Integer + Float + string)
*/

type number interface {
	constraints.Integer | constraints.Float
}

func Max[T number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func F3() {
	max := Max(45.1, 24.2)
	min := Min(45.4, 24.5)
	fmt.Printf("%v,%T\n", max, max)
	fmt.Printf("%v,%T\n", min, min)
}
