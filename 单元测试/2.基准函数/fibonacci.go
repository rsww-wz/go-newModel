package __基准函数

// 计算斐波那契数列的第n项
// 0 1 1 2 3 5 8 13 21 34 55 89
func Fibonacci(n int) int {
	a := 0
	b := 1
	res := 1

	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	for i := 2; i < n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}
