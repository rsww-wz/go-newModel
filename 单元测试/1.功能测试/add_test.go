package __功能测试

import "testing"

/*
测试组
	如果想在一个函数中测试多个案例

	可以把测试案例和结果分别用一个切片存放
	遍历切片执行即可
	但是这样不能指定执行某个案例，不够灵活

	可以是go test提供的测试组
		定义一个测试用例的结构体
		用for循环执行测试用例
		通过参数指定执行哪个测试用例

*/

func TestAdd(t *testing.T) {
	// 定义测试结构体
	type test struct {
		input [2]int
		want  int
	}

	// 定义测试案例
	tests := []test{
		{[2]int{0, 0}, 0},
		{[2]int{-1, 1}, 0},
		{[2]int{-1, -2}, -3},
		{[2]int{2, 2}, 4},
		{[2]int{3, -2}, 1},
		{[2]int{3, -3}, 0},
	}

	// 执行测试案例
	for _, v := range tests {
		got := Add(v.input[0], v.input[1])
		want := v.want

		if got != want {
			t.Errorf("want:%v got:%v", v.input, v.want)
		}
	}
}
