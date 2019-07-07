package stack

type QueueWithStack struct {
	s1 *Stack
	s2 *Stack
}

func NewQueueUsingStack() *QueueWithStack {
	return &QueueWithStack{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (q *QueueWithStack) EnQueue1(data interface{}) {
	for !q.s1.IsEmpty() {
		q.s2.Push(q.s1.Pop())
	}
	q.s1.Push(data)

	for !q.s2.IsEmpty() {
		q.s1.Push(q.s2.Pop())
	}
}

func (q *QueueWithStack) DeQueue1() interface{} {
	return q.s1.Pop()
}

func (q *QueueWithStack) EnQueue2(data interface{}) {
	q.s1.Push(data)
}

func (q *QueueWithStack) DeQueue2() interface{} {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Pop()
}

func (q *QueueWithStack) EnQueue3(data interface{}) {
	q.s1.Push(data)
}

func (q *QueueWithStack) DeQueue3() interface{} {
	if q.s1.IsEmpty() {
		return -1
	}

	x := q.s1.Pop()
	if q.s1.IsEmpty() {
		return x
	}

	item := q.DeQueue3()
	q.EnQueue3(x)
	return item
}

func (q *QueueWithStack) Size() int {
	return q.s1.Size()
}
