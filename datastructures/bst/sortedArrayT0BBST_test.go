package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedArrayToBBSTM1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	bst := SortedArrayToBBSTM1(input)
	assert.Equal(t, len(input), bst.Size())
	inorder := bst.Inorder()
	for i := range input {
		assert.Equal(t, input[i], inorder[i])
	}
}
