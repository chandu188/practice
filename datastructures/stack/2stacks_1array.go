package stack

import "errors"

type _2Stacks struct {
	data []int
	top1 int
	top2 int
}

func New2Stacks(n int) *_2Stacks {
	return &_2Stacks{
		data: make([]int, n),
		top2: n,
		top1: -1,
	}
}

func (s *_2Stacks) Push1(n int) error {
	if s.top1 < s.top2-1 {
		s.top1 = s.top1 + 1
		s.data[s.top1] = n
		return nil
	}
	return errors.New("backed array is full")
}

func (s *_2Stacks) Pop1() int {
	if s.top1 > -1 && s.top1 < len(s.data) {
		i := s.data[s.top1]
		s.top1 = s.top1 - 1
		return i
	}
	return -1
}

func (s *_2Stacks) Push2(n int) error {
	if s.top1 < s.top2-1 {
		s.top2 = s.top2 - 1
		s.data[s.top2] = n
		return nil
	}
	return errors.New("backed array is full")
}

func (s *_2Stacks) Pop2() int {
	if s.top2 > -1 && s.top2 < len(s.data) {
		i := s.data[s.top2]
		s.top2 = s.top2 + 1
		return i
	}
	return -1
}
