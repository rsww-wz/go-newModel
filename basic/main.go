package main

import (
	Stack "basic/container/stack"
	"fmt"
)

func main() {
	TestListStack()
}

func TestArrStack() {
	stack := Stack.NewArrStack[int](10)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack.Size())
	fmt.Println(stack.IsEmpty())
	v := stack.Pop()
	fmt.Printf("%v,%T\n", v, v)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

func TestListStack() {
	stack := Stack.NewListStack[string]()
	stack.Push("hello")
	stack.Push("world")
	stack.Push("Good")
	stack.Push("morning")
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
}
