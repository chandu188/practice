package bst

func AddGreaterKeysToAllTheKeys(bst *Bst)  {
	sum := 0
	addGreaterKeysToAllTheKeys(bst.root, &sum)
}

func addGreaterKeysToAllTheKeys(n *node, sum *int) {
	if n == nil {
		return
	}

	addGreaterKeysToAllTheKeys(n.right, sum)
	*sum = *sum + n.data
	n.data = *sum
	addGreaterKeysToAllTheKeys(n.left, sum)
}