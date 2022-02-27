package __基准函数

func Swap(a, b *int) {
	*a, *b = *b, *a
}
