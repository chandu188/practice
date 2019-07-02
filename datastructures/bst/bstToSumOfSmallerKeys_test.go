package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBstToSumOfSmallerKeys(t *testing.T) {
	in := []int{9, 6, 15, 3, 21}
	bst := NewBSTFrom(in)
	BstToSumOfSmallerKeys(bst)
	act := []int{3, 9, 18, 33, 54}
	inOrder := bst.Inorder()
	assert.Equal(t, len(act), len(inOrder))
	for i := range act {
		assert.Equal(t, act[i], inOrder[i])
	}
}
