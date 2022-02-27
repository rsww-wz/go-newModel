package __基准函数

import "testing"

/*
性能基准测试(压力测试函数)
	压力测试用例函数	    必须以Benchmark开头，后面是被测试函数名字，而且不能以小写字母开头
	压力测试参数	        testing.B

变量 : b.N
	它不是固定的
	它是按照一定数量增加的序列

压力测试原理
	先执行一遍函数
	如果不超过一秒，就增加b.N
	然后继续执行
	执行的时间是1秒

	简单来说，就是测试1秒内，能把你的函数跑多少次


压力测试命令行
	go test -bench=被测试函数名字

测试结果
	BenchmarkSwap-12	CPU进程数
	1000000000          执行次数
 	0.4446 ns/op        每一次操作耗费的时间，单位纳秒


参数
	-benchmem  		每次操作内存的相关数据
		0 B/op		每次操作消耗的内存
	0 allocs/op		每次操作内存申请数量

*/

func TestSwap(t *testing.T) {
	got := [2]int{4, 7}
	Swap(&got[0], &got[1])
	want := [2]int{7, 4}
	if want != got {
		t.Errorf("want:%v got:%v", want, got)
	}
}

func BenchmarkSwap(b *testing.B) {
	got := [2]int{4, 7}
	for i := 0; i < b.N; i++ {
		Swap(&got[0], &got[1])
	}
}
