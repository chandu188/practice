package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinElementSpecialStack(t *testing.T) {
	ss := NewMinElementStack()
	ss.Push(3)
	ss.Push(5)
	assert.Equal(t, 3, ss.GetMin())
	ss.Push(2)
	ss.Push(1)
	assert.Equal(t, 1, ss.GetMin())
	assert.Equal(t, 1, ss.Pop())
	assert.Equal(t, 2, ss.GetMin())
	assert.Equal(t, 2, ss.Pop())
	assert.Equal(t, 5, ss.Pop())
}
