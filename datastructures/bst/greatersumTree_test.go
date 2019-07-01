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

func TestGreaterSumTree(t *testing.T) {
	input := []int{11, 2, 29, 1, 7, 15, 40, 35}
	asc := []int{139, 137, 130, 119, 104, 75, 40, 0}
	bst := NewBSTFrom(input)
	GreaterSumTree(bst)
	data := bst.Inorder()
	assert.Equal(t, len(asc), len(data))
	for i := range data {
		assert.Equal(t, asc[i], data[i])
	}
}
