package queue

type queueNode[T any] struct {
	data T
	next *queueNode[T]
}

type listQueue[T any] struct {
	len  int
	size int // 限定队列个数时使用，没有限制的队列大小,不动size成员
	head *queueNode[T]
	last *queueNode[T]
}

func NewListQueue[T any]() *listQueue[T] {
	return &listQueue[T]{
		len:  0,
		head: nil,
		last: nil,
	}
}

// NewLimListQueue 限制长度的队列用size限制
func NewLimListQueue[T any](size int) *listQueue[T] {
	return &listQueue[T]{
		len:  0,
		size: size,
		head: nil,
		last: nil,
	}
}

func newQNode[T any](val T) *queueNode[T] {
	return &queueNode[T]{
		data: val,
		next: nil,
	}
}

func (q *listQueue[T]) IsEmpty() bool {
	return q.len == 0
}

// IsFull 如果无长度限制的队列强制使用该函数，返回false
func (q *listQueue[T]) IsFull() bool {
	if q.size == 0 {
		return false
	}

	if q.len == q.size {
		return true
	} else {
		return false
	}
}

func (q *listQueue[T]) Size() int {
	return q.len
}

func (q *listQueue[T]) Offer(val T) {
	node := newQNode[T](val)

	if q.head == nil {
		q.head = node
		q.last = node
		q.len++
		return
	}

	q.last.next = node
	q.last = node
	q.len++
}

func (q *listQueue[T]) Poll() T {
	if q.head == nil {
		var flag T
		return flag
	}

	val := q.head.data
	q.head = q.head.next
	q.len--
	return val
}

func (q *listQueue[T]) Peek() T {
	if q.head == nil {
		var flag T
		return flag
	}
	return q.head.data
}

func (q *listQueue[T]) LimitOffer(val T) {
	if q.IsFull() {
		q.Poll()
		q.Offer(val)
		return
	}

	q.Offer(val)
}
