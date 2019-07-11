package stack

import "github.com/chandu188/practice/datastructures/queue"

type stackWithQueues struct {
	q1 *queue.Queue
	q2 *queue.Queue
}

func NewStackUsingQueue() Stack {
	return &stackWithQueues{
		q1: queue.NewQueue(),
		q2: queue.NewQueue(),
	}
}

func (s *stackWithQueues) Push(data interface{}) {
	s.q1.Push(data.(int))
	for !s.q2.IsEmpty() {
		v := s.q2.Pop()
		s.q1.Push(v)
	}
	t := s.q2
	s.q2 = s.q1
	s.q1 = t
}

func (s *stackWithQueues) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stackWithQueues) Pop() interface{} {
	return s.q2.Pop()
}

func (s *stackWithQueues) Top() interface{} {
	return s.q2.Front()
}

func (s *stackWithQueues) Size() int {
	return s.q2.Size()
}
