package list

func (l *list[T]) AddFirst(val T) {
	node := newNode[T](val)
	node.next = l.head
	node.before = nil

	if l.head == nil {
		l.head = node
		l.last = node
	} else {
		l.head = node
		node.next.before = node
	}

	l.size++
}

func (l *list[T]) AddLast(val T) {
	node := newNode(val)
	if l.head == nil {
		l.last = node
		l.head = node
		l.size++
		return
	}

	node.before = l.last
	l.last.next = node
	l.last = node
	l.size++
}

func (l *list[T]) AddList(lst []T) {
	// 反正都是在最后添加，调用addLast即可
	for _, v := range lst {
		l.AddLast(v)
	}
}
