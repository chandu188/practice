package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	 {10, 5, 1, 7, 40, 50}


						 10
					   /   \
					  5     40
					 /  \      \
					1    7      50
*/

func TestPreOrderToBSTM1(t *testing.T) {
	input := []int{10, 5, 1, 7, 40, 50}
	bst := preOrderToBST(input)
	assert.Equal(t, len(input), bst.Size())
	asc := []int{1, 5, 7, 10, 40, 50}
	inorder := bst.Inorder()
	assert.Equal(t, len(asc), bst.Size())
	for i := range asc {
		assert.Equal(t, asc[i], inorder[i])
	}
}

func TestPreOrderToBSTM2(t *testing.T) {
	input := []int{10, 5, 1, 7, 40, 50}
	bst := preOrderToBSTM2(input)
	assert.Equal(t, len(input), bst.Size())
	asc := []int{1, 5, 7, 10, 40, 50}
	inorder := bst.Inorder()
	assert.Equal(t, len(asc), bst.Size())
	for i := range asc {
		assert.Equal(t, asc[i], inorder[i])
	}
}
