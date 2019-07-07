package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueUsingStack(t *testing.T) {
	q := NewQueueUsingStack()
	in := []int{1, 2, 3}
	for _, v := range in {
		q.EnQueue1(v)
	}
	assert.Equal(t, len(in), q.Size())
	for _, v := range in {
		assert.Equal(t, v, q.DeQueue1())
	}
	assert.Equal(t, 0, q.Size())

}

func TestQueueUsingStack2(t *testing.T) {
	q := NewQueueUsingStack()
	in := []int{1, 2, 3}
	for _, v := range in {
		q.EnQueue2(v)
	}
	assert.Equal(t, len(in), q.Size())
	for _, v := range in {
		assert.Equal(t, v, q.DeQueue2())
	}
	assert.Equal(t, 0, q.Size())
}
