package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(10)
	q.Push(11)
	q.Push(12)
	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 10, q.Pop())
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 11, q.Front())
}
