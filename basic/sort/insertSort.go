package sort

import "basic/container/base"

func InsertSort[T base.Number](list []T) {
	var next T
	for outer := 1; outer < len(list); outer++ {
		now := outer - 1
		next = list[outer]

		for now >= 0 && next < list[now] {
			list[now+1] = list[now]
			now--
		}
		list[now+1] = next
	}
}
