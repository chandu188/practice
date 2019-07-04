package queue

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{data: make([]int, 0)}
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Push(data int) {
	q.data = append(q.data, data)
}

func (q *Queue) Pop() int {
	if q.Size() > 0 {
		data := q.data[0]
		q.data = q.data[1:]
		return data
	}
	return -1
}

func (q *Queue) Front() int {
	return q.data[0]
}
