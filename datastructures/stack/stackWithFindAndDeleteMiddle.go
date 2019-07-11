package stack

type dLLNode struct {
	data interface{}
	prev *dLLNode
	next *dLLNode
}

func newDLLNode(data interface{}) *dLLNode {
	return &dLLNode{
		data: data,
	}
}

type stackUsingDLL struct {
	top  *dLLNode
	size int
	mid  *dLLNode
}

func (s *stackUsingDLL) Push(data interface{}) {
	node := newDLLNode(data)
	s.size = s.size + 1
	if s.top == nil {
		s.top = node
		s.mid = node

		return
	}
	if s.size%2 == 0 {
		s.mid = s.mid.prev
	}
	node.next = s.top
	s.top.prev = node
	s.top = node
}

func (s *stackUsingDLL) Pop() interface{} {
	if s.size == 0 {
		return nil
	}

	n := s.top
	s.top = n.next
	s.top.prev = nil
	if s.size%2 == 1 {
		s.mid = s.mid.next
	}
	return n.data

}

func (s *stackUsingDLL) GetMiddle() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.mid.data
}
