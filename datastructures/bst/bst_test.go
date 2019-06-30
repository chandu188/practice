package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*			8
			   	  /	  \
				3	    10
			   / \        \
			  1   6        14
				 / \      /
			    4   7    13         */

func TestBST(t *testing.T) {
	input := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	asc := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	bst := NewBSTFrom(input)
	assert.Equal(t, len(input), bst.Size())
	inorder := bst.Inorder()
	assert.Equal(t, len(asc), len(inorder))
	for i, ele := range asc {
		assert.Equal(t, ele, inorder[i])
	}
}

func NewBSTFrom(input []int) *Bst {
	bst := NewBST()
	for _, ele := range input {
		bst.Add(ele)
	}
	return bst
}

func TestDeleteBST(t *testing.T) {
	input := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	bst := NewBSTFrom(input)
	assert.Equal(t, len(input), bst.Size())
	bst.Delete(3)
	assert.Equal(t, len(input)-1, bst.Size())
	asc := []int{1, 4, 6, 7, 8, 10, 13, 14}
	inorder := bst.Inorder()
	assert.Equal(t, len(asc), len(inorder))
	for i, ele := range asc {
		assert.Equal(t, ele, inorder[i])
	}

}
