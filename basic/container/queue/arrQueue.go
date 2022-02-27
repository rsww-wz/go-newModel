package queue

import "basic/container/base/slice"

type circleQueue[T any] struct {
	maxsize int
	tail    int // 指向队尾,初始化为0
	head    int // 指向队首,初始化为0
	array   slice.Slice[T]
}

// NewArrQueue 要空出一个位置，所以与实际的size不符
// 表面申请size个，内部实际申请size+1个与实际匹配
func NewArrQueue[T any](size int) *circleQueue[T] {
	return &circleQueue[T]{
		maxsize: size + 1,
		tail:    0,
		head:    0,
		array:   make(slice.Slice[T], size+1),
	}
}

func (c *circleQueue[T]) IsFull() bool {
	return (c.tail+1)%c.maxsize == c.head
}

func (c *circleQueue[T]) IsEmpty() bool {
	return c.tail == c.head
}

func (c *circleQueue[T]) Size() int {
	return (c.tail - c.head + c.maxsize) % c.maxsize
}

// Offer 如果队列满了，就不添加新元素了，保持队列中原来的元素
func (c *circleQueue[T]) Offer(val T) {
	if c.IsFull() {
		return
	}

	c.array[c.tail] = val
	c.tail = (c.tail + 1) % c.maxsize
}

func (c *circleQueue[T]) Poll() T {
	if c.IsEmpty() {
		var flag T
		return flag
	}

	val := c.array[c.head]
	c.head = (c.head + 1) % c.maxsize
	return val
}

func (c *circleQueue[T]) Peek() T {
	if c.IsEmpty() {
		var flag T
		return flag
	}
	return c.array[c.head]
}

// LimitOffer 队列满了，更新队列元素，即出队一个元素，再入队一个元素
func (c *circleQueue[T]) LimitOffer(val T) {
	if c.IsFull() {
		c.Poll()
		c.Offer(val)
		return
	}

	c.array[c.tail] = val
	c.tail = (c.tail + 1) % c.maxsize
}
