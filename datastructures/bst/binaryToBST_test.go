package bst
import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestBinaryTreeToBST(t *testing.T) {
	input := []int{10,2,7,8,4}
	bst := btToBSTFrom(input)
	inorder := bst.Inorder()
	asc := []int{2,4,7,8,10}
	assert.Equal(t, len(asc), len(inorder))
	for i, ele := range asc {
		assert.Equal(t, ele, inorder[i])
	}
}