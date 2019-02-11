package main

import (
	"fmt"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}
func (this *Stack) Len() int {
	return this.length
}
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func main() {
	stack := New()

	stack.Push("Navid")
	stack.Push("David")
	stack.Push("Hassan")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
