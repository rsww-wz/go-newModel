package Stack

import "basic/container/base/slice"

type arrStack[T any] struct {
	maxTop int
	top    int
	size   int
	array  slice.Slice[T]
}

func NewArrStack[T any](size int) *arrStack[T] {
	return &arrStack[T]{
		maxTop: size,
		top:    -1,
		size:   0,
		array:  make(slice.Slice[T], size),
	}
}

func (s *arrStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *arrStack[T]) Size() int {
	return s.size
}

func (s *arrStack[T]) Push(val T) {
	if s.top == s.maxTop-1 {
		return
	}

	s.top++
	s.array[s.top] = val
	s.size++
}

func (s *arrStack[T]) Pop() T {
	if s.top == -1 {
		// T类型零值
		var flag T
		return flag
	}

	val := s.array[s.top]
	s.top--
	s.size--
	return val
}

func (s *arrStack[T]) Peek() T {
	if s.top == -1 {
		var flag T
		return flag
	}

	return s.array[s.top]
}
