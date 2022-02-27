package sort

import "basic/container/base"

func SelectSort[T base.Number](list []T) {
	for outer := 0; outer < len(list)-1; outer++ {
		maxIndex := outer
		for inner := outer + 1; inner < len(list); inner++ {
			if list[inner] > list[maxIndex] {
				maxIndex = inner
			}
		}
		list[outer], list[maxIndex] = list[maxIndex], list[outer]
	}
}
