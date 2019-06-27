package bst

type node struct {
	left  *node
	right *node
	data  int
	size  int
}

type Bst struct {
	root *node
}

// NewBST returns
func NewBST() *Bst {
	return &Bst{}
}

func newNode(data int) *node {
	return &node{data: data, size: 1}
}

//Size returns the size of the bst
func (b *Bst) Size() int {
	return size(b.root)
}

func size(n *node) int {
	if n == nil {
		return 0
	}
	return n.size
}

// Add adds a new element to bst
func (b *Bst) Add(data int) *node {
	b.root = add(b.root, data)
	return b.root
}

func add(n *node, data int) *node {
	if n == nil {
		return newNode(data)
	}
	if data < n.data {
		n.left = add(n.left, data)
	} else {
		n.right = add(n.right, data)
	}
	n.size = 1 + size(n.left) + size(n.right)
	return n
}

//Search searches for the specified data in the bst
func (b *Bst) Search(data int) *node {
	return search(b.root, data)
}

func search(n *node, data int) *node {
	if n == nil {
		return nil
	}
	if n.data == data {
		return n
	}

	if data < n.data {
		return search(n.left, data)
	}

	return search(n.right, data)
}

func deleteMin(n *node) *node {
	if n.left == nil {
		return n.right
	}
	n.left = deleteMin(n.left)
	n.size = size(n.left) + size(n.right) + 1
	return n
}

func deleteMax(n *node) *node {
	if n.right == nil {
		return n.left
	}
	n.right = deleteMax(n.right)
	n.size = size(n.left) + size(n.right) + 1
	return n
}

func (b *Bst) Delete(data int) {
	b.root = delete(b.root, data)
}

func delete(n *node, data int) *node {
	if n == nil {
		return nil
	}
	if data < n.data {
		n.left = delete(n.left, data)
	} else if data > n.data {
		n.right = delete(n.right, data)
	} else {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}

		t := n
		n = min(t.right)
		n.right = deleteMin(t.right)
		n.left = t.left
	}
	n.size = size(n.left) + size(n.right) + 1
	return n
}

func min(n *node) *node {
	if n.left == nil {
		return n
	}
	return min(n.left)
}

// Inorder returns the inorder iteration of bst
func (b *Bst) Inorder() []int {
	elems := make([]int, 0)
	elems = inorder(make([]int, 0), b.root)
	return elems
}

func inorder(elems []int, n *node) []int {
	if n == nil {
		return elems
	}
	elems = inorder(elems, n.left)
	elems = append(elems, n.data)
	elems = inorder(elems, n.right)
	return elems
}
