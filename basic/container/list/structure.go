package list

/*双循环链表*/

type listNode[T any] struct {
	data   T
	before *listNode[T]
	next   *listNode[T]
}

type list[T any] struct {
	head *listNode[T]
	last *listNode[T]
	size int
}

func newNode[T any](data T) *listNode[T] {
	return &listNode[T]{
		data:   data,
		before: nil,
		next:   nil,
	}
}

func NewList[T any]() *list[T] {
	return &list[T]{
		head: nil,
		last: nil,
		size: 0,
	}
}
