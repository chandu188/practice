package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePathBSTUsingQueue(t *testing.T) {
	in := []int{50, 30, 70, 20, 40, 60, 80}
	bst := NewBSTFrom(in)
	ReversePathFromNode(bst, 70)
	act := bst.Inorder()
	exp := []int{20, 30, 40, 70, 60, 50, 80}
	assert.Equal(t, len(exp), len(act))
	for i := range exp {
		assert.Equal(t, exp[i], act[i])
	}
}
