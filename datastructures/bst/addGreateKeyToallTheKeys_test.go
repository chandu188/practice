package bst

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestAddGreaterKeysToAllTheKeys(t *testing.T){
	input := []int{5,2, 13}
	bst := NewBSTFrom(input)
	AddGreaterKeysToAllTheKeys(bst)
	act := []int{20, 18, 13}
	inOrder := bst.Inorder()
	assert.Equal(t, len(act), len(inOrder))
	for i := range act {
		assert.Equal(t, act[i], inOrder[i])
	}
}