package DataStructure

type stack struct {
	value interface{}
	next  *stack
}

func (s stack) Pull() stack {
	return *s.next
}

func (s stack) Push(stackValue interface{}) {
	s = stack{value: stackValue, next: &s}
}
