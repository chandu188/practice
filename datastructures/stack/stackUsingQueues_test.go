package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackUsingQueue(t *testing.T) {
	st := NewStackUsingQueue()
	assert.Equal(t, true, st.IsEmpty())
	for _, v := range []int{1, 2, 3, 4} {
		st.Push(v)
	}

	assert.Equal(t, 4, st.Size())
	assert.Equal(t, 4, st.Top())
	assert.Equal(t, 4, st.Pop())
	assert.Equal(t, 3, st.Pop())
	assert.Equal(t, 2, st.Size())
}
