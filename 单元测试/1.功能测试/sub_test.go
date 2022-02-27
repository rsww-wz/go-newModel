package __功能测试

import "testing"

/*
子测试
	go1.7中加入了子测试,t.run
	for循环中，用t.run执行每个测试案例
	go test -v 可以看到每个测试案例的具体情况
	也可以执行某个单独的子测试

	t.Run(name string, func(t *testing.T){})
		里面是个匿名函数

	但是要求每个测试案例都要有名字
	所以可以用map存放测试案例


单独执行某个测试用例
	-run=被试函数名/测试用例的名字
	后面不用指定测试函数了

其他参数
	-cover 测试覆盖率
		检测被测试函数覆盖比率
		也就是代码能跑到的行数
		一般至少要60%以上

	-coverprofile=文件
		把测试的具体情况导出到指定文件中
		go tool cover -html=文件
		可以以网页的形式打开文件查看具体情况

注意：=前后不能有空格
*/

func TestSub(t *testing.T) {
	// 定义测试结构体
	type test struct {
		input [2]int
		want  int
	}

	// 定义测试案例
	tests := map[string]test{
		"1": {[2]int{0, 0}, 0},
		"2": {[2]int{-1, 1}, -2},
		"3": {[2]int{-1, -2}, 1},
		"4": {[2]int{2, 2}, 0},
		"5": {[2]int{3, -2}, 5},
		"6": {[2]int{3, -3}, 6},
	}

	// 执行测试案例
	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			got := Sub(v.input)

			if got != v.want {
				t.Errorf("name:%s failed, want:%v, got:%v", k, v.input, v.want)
			}
		})
	}
}
