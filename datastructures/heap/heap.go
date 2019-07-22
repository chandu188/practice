package heap

type MinPQ struct {
	data  []int
	index int
}

func (pq *MinPQ) Size() int {
	return len(pq.data) - 1
}

func Heapify(data []int) *MinPQ {
	pq := NewMinPQ(data)
	for k := pq.Size() / 2; k >= 1; k-- {
		pq.Sink(k)
	}
	return pq
}

func NewMinPQ(data []int) *MinPQ {
	d := make([]int, len(data)+1)
	for i, v := range data {
		d[i+1] = v
	}
	return &MinPQ{
		data:  d,
		index: len(d) - 1,
	}
}

func swap(data []int, i, j int) {
	temp := data[i]
	data[i] = data[j]
	data[j] = temp
}

func less(data []int, i, j int) bool {
	if data[i] < data[j] {
		return true
	}
	return false
}

func (pq *MinPQ) Sink(k int) {
	data := pq.data
	n := pq.index
	for 2*k <= n {
		j := 2 * k
		if j < n && less(data, j+1, j) {
			j++
		}
		if less(data, k, j) {
			break
		}
		swap(data, k, j)
		k = j
	}
}

func (pq *MinPQ) Swim(k int) {
	data := pq.data
	for k > 1 && less(data, k, k/2) {
		swap(data, k, k/2)
		k = k / 2
	}
}

func (pq *MinPQ) Add(val int) {
	pq.index++
	pq.data[pq.index] = val
	pq.Swim(pq.index)
}

func (pq *MinPQ) DelMin() int {
	min := pq.data[1]
	swap(pq.data, 1, pq.index)
	pq.index--
	pq.Sink(1)
	return min
}
