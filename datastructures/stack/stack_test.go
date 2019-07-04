package stack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func  TestStack(t *testing.T)  {
	s := NewStack()
	data := []int{3,7,6,8}
	for _, v := range data {
		s.Push(v)
	}
		
	assert.Equal(t, len(data), s.Size())
	assert.Equal(t, 8, s.Pop())
	assert.Equal(t, len(data)-1, s.Size())

}