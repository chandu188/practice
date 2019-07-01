package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTerenaryTrie(t *testing.T) {
	tt := &TerenaryTrie{}
	input := []struct {
		key   string
		value int
	}{
		{
			key:   "she",
			value: 10,
		},
		{
			key:   "sells",
			value: 1,
		},
		{
			key:   "sea",
			value: 2,
		},
		{
			key:   "shells",
			value: 3,
		},
		{
			key:   "by",
			value: 4,
		},
		{
			key:   "the",
			value: 5,
		},
		{
			key:   "sea",
			value: 6,
		},
		{
			key:   "shore",
			value: 7,
		},
	}
	for _, in := range input {
		tt.Put(in.key, in.value)
	}

	testCases := []struct {
		key   string
		value int
	}{
		{key: "sea", value: 6}, {key: "shelter", value: -1},
	}

	for _, tc := range testCases {
		act := tt.Get(tc.key)
		assert.Equal(t, tc.value, act)
	}
}

func TestTerenaryTrieKeys(t *testing.T) {
	tt := &TerenaryTrie{}
	input := []struct {
		key   string
		value int
	}{
		{
			key:   "she",
			value: 0,
		},
		{
			key:   "sells",
			value: 1,
		},
		{
			key:   "sea",
			value: 2,
		},
		{
			key:   "shells",
			value: 3,
		},
		{
			key:   "by",
			value: 4,
		},
		{
			key:   "the",
			value: 5,
		}, {
			key:   "sea",
			value: 6,
		}, {
			key:   "shore",
			value: 7,
		},
	}
	for _, in := range input {
		tt.Put(in.key, in.value)
	}
	res := []string{"she", "sea", "shells", "by", "the", "shore"}
	keys := tt.Keys()
	assert.Equal(t, len(res), len(keys))
}

func TestTerenaryTrieKeysWithPrefix(t *testing.T) {
	tt := &TerenaryTrie{}
	input := []struct {
		key   string
		value int
	}{
		{
			key:   "she",
			value: 10,
		},
		{
			key:   "sells",
			value: 1,
		},
		{
			key:   "sea",
			value: 2,
		},
		{
			key:   "shells",
			value: 3,
		},
		{
			key:   "by",
			value: 4,
		},
		{
			key:   "the",
			value: 5,
		}, {
			key:   "sea",
			value: 6,
		},
		{
			key:   "shore",
			value: 7,
		},
	}
	for _, in := range input {
		tt.Put(in.key, in.value)
	}
	res := []string{"she", "shells", "shore"}
	keys := tt.KeysWithPrefix("sh")
	assert.Equal(t, len(res), len(keys))
}

func TestLongestKey(t *testing.T) {
	tt := &TerenaryTrie{}
	input := []struct {
		key   string
		value int
	}{
		{
			key:   "she",
			value: 10,
		},
		{
			key:   "sells",
			value: 1,
		},
		{
			key:   "sea",
			value: 2,
		},
		{
			key:   "shells",
			value: 3,
		},
		{
			key:   "by",
			value: 4,
		},
		{
			key:   "the",
			value: 5,
		}, {
			key:   "sea",
			value: 6,
		},
		{
			key:   "shore",
			value: 7,
		},
	}
	for _, in := range input {
		tt.Put(in.key, in.value)
	}
	longKey := tt.LongestPrefixOf("shell")
	assert.Equal(t, "shell", longKey)
	longKey = tt.LongestPrefixOf("she")
	assert.Equal(t, "she", longKey)
	longKey = tt.LongestPrefixOf("shellsort")
	assert.Equal(t, "shells", longKey)
	longKey = tt.LongestPrefixOf("by")
	assert.Equal(t, "by", longKey)
}
