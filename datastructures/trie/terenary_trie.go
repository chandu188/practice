package trie

type terenaryTrieNode struct {
	value  int
	ch     byte
	left   *terenaryTrieNode
	right  *terenaryTrieNode
	middle *terenaryTrieNode
}

type TerenaryTrie struct {
	root *terenaryTrieNode
}

func (tt *TerenaryTrie) Get(key string) int {
	n := getValue(tt.root, key, 0)
	if n == nil {
		return -1
	}
	return n.value
}

func getValue(tn *terenaryTrieNode, key string, index int) *terenaryTrieNode {
	if tn == nil {
		return nil
	}
	ch := key[index]
	if ch < tn.ch {
		return getValue(tn.left, key, index)
	} else if ch > tn.ch {
		return getValue(tn.right, key, index)
	} else if index < len(key)-1 {
		return getValue(tn.middle, key, index+1)
	} else {
		return tn
	}
}

func (tt *TerenaryTrie) Put(key string, value int) {
	tt.root = putValue(tt.root, key, value, 0)
}

func putValue(tn *terenaryTrieNode, key string, value int, index int) *terenaryTrieNode {
	ch := key[index]
	if tn == nil {
		tn = &terenaryTrieNode{}
		tn.ch = ch
	}
	if ch < tn.ch {
		tn.left = putValue(tn.left, key, value, index)
	} else if ch > tn.ch {
		tn.right = putValue(tn.right, key, value, index)
	} else if index < len(key)-1 {
		tn.middle = putValue(tn.middle, key, value, index+1)
	} else {
		tn.value = value
	}
	return tn
}

func (tt *TerenaryTrie) Keys() []string {
	keys := make([]string, 0)
	return collect(tt.root, "", keys)
}

func (tt *TerenaryTrie) KeysWithPrefix(str string) []string {
	n := getValue(tt.root, str, 0)
	keys := make([]string, 0)
	return collect(n.middle, "", keys)
}

func (tt *TerenaryTrie) LongestPrefixOf(key string) string {
	len := searchFor(tt.root, key, 0, 0)
	return key[0:len]
}
func searchFor(tn *terenaryTrieNode, key string, index int, length int) int {
	if tn == nil {
		return length
	}
	if tn.value != 0 {
		length = index + 1
	}
	if index == len(key)-1 {
		return len(key)
	}
	ch := key[index]
	if ch < tn.ch {
		return searchFor(tn.left, key, index, length)
	} else if ch > tn.ch {
		return searchFor(tn.right, key, index, length)
	} else {
		return searchFor(tn.middle, key, index+1, length)
	}
}

func collect(tt *terenaryTrieNode, prefix string, keys []string) []string {
	if tt == nil {
		return keys
	}

	keys = collect(tt.left, prefix, keys)
	keys = collect(tt.right, prefix, keys)
	keys = collect(tt.middle, prefix+string(tt.ch), keys)

	if tt.value != 0 {
		keys = append(keys, prefix)
	}
	return keys
}
