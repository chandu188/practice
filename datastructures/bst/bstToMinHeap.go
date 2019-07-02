package bst

func BSTToMinHeap(bst *Bst){
	inOrder := bst.Inorder()
	index := 0
	bstToMinHeap1(bst.root, inOrder, &index)
}

func bstToMinHeap1(n *node, data []int, index *int) *node{
	if n== nil {
		return nil
	}

	n.data = data[*index]
	*index = *index +1
	bstToMinHeap1(n.left, data, index)
	bstToMinHeap1(n.right, data, index)
	return n
}