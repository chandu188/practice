package bst

func LevelOrderToBST(data []int) *Bst {
	var root *node
	for _, v := range data {
		root = levelOrderToBST(root, v)
	}
	return &Bst{root: root}
}

func levelOrderToBST(n *node, data int) *node {
	if n == nil {
		return newNode(data)
	}
	if data <= n.data {
		n.left = levelOrderToBST(n.left, data)
	} else {
		n.right = levelOrderToBST(n.right, data)
	}
	return n
}
