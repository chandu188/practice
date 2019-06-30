package trie

type terenaryTrieNode struct{
	value int
	ch byte
	left *terenaryTrieNode
	right *terenaryTrieNode
	middle *terenaryTrieNode
}

type TerenaryTrie struct {
	root *terenaryTrieNode
}

func (tt *TerenaryTrie) Get(key string) int {
	return getValue(tt.root, key, 0)
}

func getValue(tn *terenaryTrieNode, key string, index int) int {
	if tn == nil {
		return -1
	} 
	
	if index == len(key) -1 {
		return tn.value
	}

	if tn.ch == key[index] {
		return getValue(tn.middle, key, index + 1)
	} else if key[index] < tn.ch {
		return getValue(tn.left, key, index + 1)
	}
	return getValue(tn.right, key, index + 1)
}

func (tt *TerenaryTrie) Put(key string, value int ) {
	tt.root = putValue(tt.root, key, value, 0)
}

func newTerenaryTrieNode() *terenaryTrieNode {
	return &terenaryTrieNode{}
}

func putValue(tn *terenaryTrieNode, key string, value int, index int) *terenaryTrieNode{
	if tn == nil {
		tn = newTerenaryTrieNode()
	}
	if index == len(key)-1{
		tn.value = value
		tn.ch = key[index]
		return tn
	}
	if tn.ch == key[index] {
		tn.middle = putValue(tn.middle, key, value, index+1)
	}else if key[index] < tn.ch {
		tn.left = putValue(tn.left, key, value, index +1)
	} else {
		tn.right = putValue(tn.right, key, value, index + 1)
	}
	return tn
	}