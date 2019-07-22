package bst

const RED = true
const BLACK = false

type rbNode struct {
	left  *rbNode
	right *rbNode
	color bool
}

type rbTree struct {
	root *rbNode
}

func rotateLeft(n *rbNode) *rbNode {
	x := n.right
	n.right = x.left
	x.left = n
	x.color = n.color
	n.color = RED
	return x
}

func rotateRight(n *rbNode) *rbNode {
	x := n.left
	n.left = x.right
	x.right = n
	x.color = n.color
	n.color = RED
	return x
}

func flipColors(n *rbNode) {
	n.left.color = BLACK
	n.right.color = BLACK
	n.color = RED
}
