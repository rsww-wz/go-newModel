package list

func (l *list[T]) DeleteFirst() T {
	var res T

	// 链表为空
	if l.head == nil {
		var flag T
		return flag
	}

	res = l.head.data

	// 只有一个元素
	if l.head.next == nil {
		l.head = nil
		l.last = nil
		l.size--
		return res
	}

	// 多个元素
	l.head = l.head.next
	l.head.before = nil
	l.size--
	return res
}

func (l *list[T]) DeleteLast() T {
	var res interface{}

	// 链表为空
	if l.head == nil {
		var flag T
		return flag
	}

	// 只有一个元素
	if l.head.next == nil {
		res = l.head.data
		l.head = nil
		l.last = nil
		l.size--
		return res
	}

	// 多个元素
	res = l.last.data
	l.last = l.last.before
	l.last.next.before = nil
	l.last.next = nil
	l.size--
	return res
}

func (l *list[T]) RemoveFirst() {
	l.DeleteFirst()
}

func (l *list[T]) RemoveLast() {
	l.DeleteLast()
}
