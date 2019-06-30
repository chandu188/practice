package trie

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestTrie(t *testing.T) {
	input := []struct{
		key string
		value int
	}{
		{
			key: "by",
			value: 4,
		},
		{
			key: "Shell",
			value: 42,
		},
		{
			key: "Sort",
			value: 41,
		},
		{
			key: "MergeSort",
			value: 2,
		},
	}

	tr := trie{}
	for _, n := range input {
		tr.Put(n.key,n.value)
	}

	assert.Equal(t, 4, tr.Get("by"))
	assert.Equal(t, 42, tr.Get("Shell"))
	assert.Equal(t, 41, tr.Get("Sort"))
	assert.Equal(t, 2, tr.Get("MergeSort"))

}