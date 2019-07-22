package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortUsingStack(t *testing.T) {
	st := NewStack()
	data := []int{30, -5, 18, 14, -3}
	for _, v := range data {
		st.Push(v)
	}
	SortStack(st)
	exp := []int{30, 18, 14, -3, -5}
	for _, v := range exp {
		assert.Equal(t, v, st.Pop())
	}

}
