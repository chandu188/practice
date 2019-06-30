package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	data := []int{1,2,3}
	ll := newLLFrom(data)
	assert.Equal(t, len(data), ll.size)
	//1->2->3->4->5->6->7
	data = []int{1,2,3,4,5,6,7}
	ll = newLLFrom(data)
	assert.Equal(t, len(data), ll.size)
	assert.Equal(t, ll.front.data, 1)
	assert.Equal(t, ll.rear.data, 7)
}

func TestSortedLLToBBST(t *testing.T) {
	data := []int{1,2,3,4,5,6,7}
	ll := newLLFrom(data)
	bst := SortedLLTOBBST(ll)
	inorder := bst.Inorder()
	assert.Equal(t, len(data), len(inorder))
	for i, ele := range data {
		assert.Equal(t, ele, inorder[i])
	}

}