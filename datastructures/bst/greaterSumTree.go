package bst

func GreaterSumTree(t *Bst) *Bst {
	sum := 0
	greaterSumTree(t.root, &sum)
	return t
}

func greaterSumTree(n *node, sum *int) {
	if n == nil {
		return
	}
	greaterSumTree(n.right, sum)
	*sum = *sum + n.data
	n.data = *sum - n.data
	greaterSumTree(n.left, sum)
}
