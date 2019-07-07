package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseStackRecursion(t *testing.T) {
	st := NewStack()
	data := []int{1, 2, 3, 4}
	for _, v := range data {
		st.Push(v)
	}
	ReverseStack(st)
	exp := []int{1, 2, 3, 4}
	for _, v := range exp {
		assert.Equal(t, v, st.Pop())
	}
}
