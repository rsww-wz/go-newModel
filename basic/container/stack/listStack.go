package Stack

type stackNode[T any] struct {
	data T
	next *stackNode[T]
}

type listStack[T any] struct {
	len  int
	root *stackNode[T]
}

func NewListStack[T any]() *listStack[T] {
	return &listStack[T]{
		len:  0,
		root: nil,
	}
}

func newStackNode[T any](val T) *stackNode[T] {
	return &stackNode[T]{
		data: val,
		next: nil,
	}
}

func (s *listStack[T]) IsEmpty() bool {
	return s.len == 0
}

func (s *listStack[T]) Size() int {
	return s.len
}

func (s *listStack[T]) Push(val T) {
	node := newStackNode[T](val)
	if s.root == nil {
		s.root = node
		s.len++
		return
	}

	node.next = s.root
	s.root = node
	s.len++
}

func (s *listStack[T]) Pop() T {
	var flag T
	if s.root == nil {
		return flag
	}

	val := s.root.data
	s.root = s.root.next
	s.len--
	return val
}

func (s *listStack[T]) Peek() T {
	var flag T
	if s.root == nil {
		return flag
	}
	return s.root.data
}
