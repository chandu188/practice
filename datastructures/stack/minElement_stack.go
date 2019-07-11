package stack

type SpecialStack struct {
	minElement int
	Stack
}

func NewMinElementStack() *SpecialStack {
	return &SpecialStack{Stack: NewStack()}
}

func (s *SpecialStack) Push(elem int) {
	if s.Stack.IsEmpty() {
		s.Stack.Push(elem)
		s.minElement = elem
		return
	}

	if elem > s.minElement {
		s.Stack.Push(elem)
		return
	}

	prevMinElement := s.minElement
	s.minElement = elem
	elem = 2*elem - prevMinElement
	s.Stack.Push(elem)
}

func (s *SpecialStack) GetMin() int {
	return s.minElement
}

func (s *SpecialStack) Pop() int {
	ele := (s.Stack.Pop()).(int)
	if ele > s.minElement {
		return ele
	}
	v := s.minElement
	s.minElement = 2*v - ele
	return v

}
