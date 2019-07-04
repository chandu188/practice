package trie

type trieNode struct {
	value int
	keys  [123]*trieNode
}

type trie struct {
	root *trieNode
}

func (t *trie) Put(key string, value int) {
	t.root = put(t.root, key, value, 0)
}

func newTrieNode() *trieNode {
	return &trieNode{}
}

func put(n *trieNode, key string, value int, index int) *trieNode {
	if n == nil {
		n = newTrieNode()
	}
	if len(key)-1 == index {
		n.value = value
		return n
	}
	ch := key[index]
	n.keys[ch] = put(n.keys[ch], key, value, index+1)

	return n
}

func (t *trie) Get(key string) int {
	return get(t.root, key, 0)
}

func get(n *trieNode, key string, index int) int {
	if n == nil {
		return -1
	}
	if len(key)-1 == index {
		return n.value
	}
	ch := key[index]
	n = n.keys[ch]
	return get(n, key, index+1)
}
