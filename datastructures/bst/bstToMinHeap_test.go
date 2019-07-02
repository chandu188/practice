package bst

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestBSTToMinHeap(t *testing.T){
	in := []int{4,2,6,1,3,5,7}
	bst := NewBSTFrom(in)
	BSTToMinHeap(bst)
	inorder := bst.Inorder()
	assert.Equal(t, len(in), len(inorder))
	act := []int{3,2,4,1,6,5,7}
	for i := range inorder{
		assert.Equal(t, act[i], inorder[i])
	}

}