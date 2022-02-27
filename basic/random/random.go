package random

import (
	"constraints"
	"math/rand"
	"time"
)

func Random() {
	rand.Seed(time.Now().UnixNano())
}

// List 生成指定大小的有序切片，范围0-length
func List[T constraints.Integer](length T) []T {
	list := make([]T, length, length)
	var i T
	for i = 0; i < length; i++ {
		list[i] = i
	}
	return list
}

func RangeInt[T constraints.Integer](start, end T) []T {
	var (
		size = end - start
		list = make([]T, size, size)
		i    T
	)

	for i = 0; i < size; i++ {
		list[i] = start + i
	}

	return list
}

func Shuffle[T constraints.Integer](list []T) []T {
	var i = len(list) - 1
	var j int

	for ; i >= 0; i-- {
		j = rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}

	return list
}
