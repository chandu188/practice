package stack

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Pop() interface{} {
	if s.Size() > 0 {
		data := s.data[len(s.data)-1]
		s.data = s.data[:s.Size()-1]
		return data
	}

	return -1

}
