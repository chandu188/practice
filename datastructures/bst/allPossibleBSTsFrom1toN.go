package bst

func GenerateBSTs(upTo int) []*Bst {
	roots := generateBst(1, upTo)
	bsts := make([]*Bst, 0)
	for _, root := range roots {
		bsts = append(bsts, &Bst{root: root})
	}
	return bsts

}

func generateBst(start int, end int) []*node {
	roots := make([]*node, 0)
	if start > end {
		roots = append(roots, nil)
		return roots
	}

	for i := start; i <= end; i++ {
		leftRoots := generateBst(start, i-1)

		rightRoots := generateBst(i+1, end)
		for j := 0; j < len(leftRoots); j++ {
			left := leftRoots[j]
			for k := 0; k < len(rightRoots); k++ {
				right := rightRoots[k]
				node := newNode(i)
				node.left = left
				node.right = right
				roots = append(roots, node)
			}

		}
	}

	return roots

}
