package bst

func BSTToMinHeapInPlace(bst *Bst) {
	var head *node
	head = nil
	bstToSLL(bst.root, &head)
	var root *node
	root = nil
	root = sllToMinHeap(root, head)
	bst.root = root
}

func bstToSLL(n *node, head **node) *node {
	if n == nil {
		return nil
	}
	bstToSLL(n.right, head)
	n.right = *head
	if *head != nil {
		(*head).left = nil
	}
	*head = n
	bstToSLL(n.left, head)
	return n
}

func sllToMinHeap(n *node, head *node) *node{
	if head== nil {
		return nil
	}

	var root *node
	root= head
	head = head.right

	q := make([]*node, 0)
	q = append(q, root)

	for ;head!=nil;{
		parent := q[0]
		q = q[1:]

		left := head
		head = head.right
		left.right = nil
		q = append(q, left)
		parent.left = left

		if head != nil {
			right := head
			head = head.right
			right.right = nil
			q = append(q, right)
			parent.right = right
		}
	}
	return root
}

