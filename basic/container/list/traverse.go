package list

import "fmt"

func (l *list[T]) Traverse() {
	if l.head == nil {
		fmt.Println()
		return
	}

	// todo bug
	for temp := l.head; temp != nil; temp = temp.next {
		fmt.Print(temp.data, "==>")
	}
	fmt.Println()
}

func (l *list[T]) ReverseTraverse() {
	if l.head == nil {
		fmt.Println()
		return
	}

	// todo bug
	for temp := l.last; temp != nil; temp = temp.before {
		fmt.Print(temp.data, "==>")
	}

	fmt.Println()
}
