package bst

// SortedArrayToBBSTM1
func SortedArrayToBBSTM1(data []int) *Bst {
	index := 0
	return &Bst{root: sortedArrayToBBSTM1(data, len(data), &index)}
}

func sortedArrayToBBSTM1(data []int, n int, index *int) *node {
	if n <= 0 {
		return nil
	}

	left := sortedArrayToBBSTM1(data, n/2, index)

	node := newNode(data[*index])
	*index = *index + 1
	node.left = left
	node.right = sortedArrayToBBSTM1(data, n-n/2-1, index)
	node.size = size(node.left) + size(node.right) + 1
	return node
}
