package grammar

import (
	"fmt"
	"strconv"
)

/*
接口的约束
	接口中写类型就是类型约束
	接口中写方法就是方法约束

	他们还可以组合使用
	如果用一个接口约束，只有是接口里面的类型或者实现了接口里面的方法才能使用泛型
*/

type myInt int

func (m myInt) string() string {
	return strconv.Itoa(int(m))
}

type stringer interface {
	string() string
}

// 约束实现了stringer接口的类型
func toString[T stringer](s []T) []string {
	var r []string
	for _, v := range s {
		r = append(r, v.string())
	}
	return r
}

func F5() {
	lst := []myInt{34, 126, 735}
	res := toString(lst)
	fmt.Printf("%v,%T", res, res[0])
}
