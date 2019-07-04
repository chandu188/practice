package bst

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrderToBST(t *testing.T) {
	in := []int{7, 4, 12, 3, 6, 8, 1, 5, 10}
	bst := LevelOrderToBST(in)
	inOrder := bst.Inorder()
	act := []int{1,3,4,5,6,7,8,10,12}
	assert.Equal(t, len(act), len(inOrder))
	for i := range act {
		assert.Equal(t, act[i], inOrder[i])
	}

}