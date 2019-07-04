package bst

import "github.com/chandu188/practice/datastructures/queue"

func ReversePathFromNode(bst *Bst, key int) {
	q := queue.NewQueue()
	bst.root = reversePathFromNode(bst.root, key, q)
}

func reversePathFromNode(n *node, key int, q *queue.Queue) *node {
	if n == nil {
		return nil
	}

	if n.data == key {
		q.Push(n.data)
		n.data = q.Pop()
		return n
	}

	if key < n.data {
		q.Push(n.data)
		reversePathFromNode(n.left, key, q)
		n.data = q.Pop()
	} else {
		q.Push(n.data)
		reversePathFromNode(n.right, key, q)
		n.data = q.Pop()
	}
	return n
}
