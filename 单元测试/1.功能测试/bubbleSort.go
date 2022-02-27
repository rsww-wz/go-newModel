package __功能测试

// 待测试函数
func BubbleSort(lst []int) {
	for i := 0; i < len(lst)-1; i++ {
		for j := 0; j < len(lst)-1-i; j++ {
			if lst[j] > lst[j+1] {
				lst[j], lst[j+1] = lst[j+1], lst[j]
			}
		}
	}
}
