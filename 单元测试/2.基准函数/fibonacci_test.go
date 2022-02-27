package __基准函数

import "testing"

func TestFibonacci(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	tests := []test{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 5},
		{7, 8},
		{8, 13},
		{12, 89},
	}

	for _, v := range tests {
		got := Fibonacci(v.input)

		if got != v.want {
			t.Errorf("want:%v got:%v", v.want, got)
		}
	}
}

// 性能测试
// 函数的性能随着参数变化
// 当传入的参数不一样，执行的次数也会不一样
// 这时候可以用性能比较函数
// 性能比较函数通常是一个带有参数的函数，被多个不同的benchmark函数传入不同的值来调用
// -bench=.  用通配符可以全部执行

func benchmarkFibonacci(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		Fibonacci(size)
	}
}

func BenchmarkFibonacci1(b *testing.B)  { benchmarkFibonacci(b, 1) }
func BenchmarkFibonacci10(b *testing.B) { benchmarkFibonacci(b, 10) }
func BenchmarkFibonacci20(b *testing.B) { benchmarkFibonacci(b, 20) }
func BenchmarkFibonacci40(b *testing.B) { benchmarkFibonacci(b, 40) }
