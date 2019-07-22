package heap

import "testing"
import "github.com/stretchr/testify/assert"

func TestMinHeap(t *testing.T) {
	in := []int{6, 3, 2, 8, 9, 0, 10}
	data := make([]int, len(in)+1)
	minHeap := &MinPQ{
		data: data,
	}
	for _, v := range in {
		minHeap.Add(v)
	}
	res := []int{0, 2, 3, 6, 8, 9, 10}
	for i := 0; i < len(in); i++ {
		assert.Equal(t, res[i], minHeap.DelMin())
	}

	assert.Equal(t, len(in), minHeap.Size())
}

func TestHeapify(t *testing.T) {
	in := []int{6, 3, 2, 8, 9, 0, 10}
	minPq := Heapify(in)
	res := []int{0, 2, 3, 6, 8, 9, 10}
	for i := 0; i < len(in); i++ {
		assert.Equal(t, res[i], minPq.DelMin())
	}

	assert.Equal(t, len(in), minPq.Size())
}
