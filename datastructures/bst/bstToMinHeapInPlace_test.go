package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*8
/    \
4      12
/  \     /  \
2    6   10  14 */
func TestBSTToMinHeapInPlace(t *testing.T) {
	in := []int{8, 4, 12, 2, 6, 10, 14}
	bst := NewBSTFrom(in)
	BSTToMinHeapInPlace(bst)
	inOrder := bst.Inorder()
	act := []int{8, 4, 10, 2, 12, 6, 14}
	assert.Equal(t, len(in), len(inOrder))
	for i := range inOrder {
		assert.Equal(t, act[i], inOrder[i])
	}
}
