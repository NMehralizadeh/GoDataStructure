package DataStructure

type Stack struct {
	value interface{}
	next  *Stack
}

func (s Stack) Pull() Stack {
	return *s.next
}

func (s Stack) Push(stackValue interface{}) {
	s = Stack{value: stackValue, next: &s}
}
