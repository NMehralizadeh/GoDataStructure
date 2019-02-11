package main

import (
	"fmt"
)

type Stack struct {
	Value interface{}
	Next  *Stack
}

func (s *Stack) New(value interface{}) {
	s = &Stack{
		Value: value,
	}
}

func (s *Stack) Pull() interface{} {
	value := s.Value
	s = s.Next
	return value
}

func (s *Stack) Push(stackValue interface{}) {
	s = &Stack{Value: stackValue, Next: s}
	fmt.Println(s)
}

func main() {
	stack := Stack{}

	stack.New("Navid")

	fmt.Println("PPPPP")
	p := &stack
	fmt.Println(*p)
	fmt.Println(stack)

	stack.Push("David")
	stack.Push("Hassan")

	fmt.Println(stack)
	fmt.Println(stack.Pull())
	fmt.Println(stack.Pull())
	fmt.Println(stack.Pull())
}
