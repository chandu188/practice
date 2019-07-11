package stack

type Stack interface {
	Push(interface{})
	Size() int
	IsEmpty() bool
	Pop() interface{}
	Top() interface{}
}

type stack struct {
	data []interface{}
}

func NewStack() Stack {
	return &stack{data: make([]interface{}, 0)}
}

func (s *stack) Push(data interface{}) {
	s.data = append(s.data, data)
}

func (s *stack) Size() int {
	return len(s.data)
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack) Pop() interface{} {
	if s.Size() > 0 {
		data := s.data[len(s.data)-1]
		s.data = s.data[:s.Size()-1]
		return data
	}

	return -1

}

func (s *stack) Top() interface{} {
	return s.data[s.Size()-1]
}
