package bst

func BstToSumOfSmallerKeys(bst *Bst) {
	sum := 0
	bstToSumOfSmallerKeys(bst.root, &sum)
}

func bstToSumOfSmallerKeys(n *node, sum *int) {
	if n == nil {
		return
	}

	bstToSumOfSmallerKeys(n.left, sum)
	*sum = *sum + n.data
	n.data = *sum
	bstToSumOfSmallerKeys(n.right, sum)

}
