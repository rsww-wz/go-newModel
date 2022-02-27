package sort

import "basic/container/base"

func BubbleSort[T base.Number](list []T) {
	for outer := 0; outer < len(list)-1; outer++ {
		for inner := 0; inner < len(list)-1-outer; inner++ {
			if list[inner] > list[inner+1] {
				list[inner], list[inner+1] = list[inner+1], list[inner]
			}
		}
	}
}
