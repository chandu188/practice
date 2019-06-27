package bst

/*
	 {10, 5, 1, 7, 40, 50}


						 10
					   /   \
					  5     40
					 /  \      \
					1    7      50
*/

func preOrderToBST(data []int) *Bst {
	n := method1(data, 0, len(data)-1)
	return &Bst{root: n}
}

func preOrderToBSTM2(data []int) *Bst {
	i := 0
	n := method2(data, &i, -1000, 10000)
	return &Bst{root: n}
}



func method1(data []int, low int, high int) *node {
	if low > high {
		return nil
	}

	if low == high {
		return newNode(data[low])
	}

	d := data[low]
	n := newNode(d)
	i := low
	for ; i <= high; i++ {
		if data[i] > d {
			break
		}
	}

	n.left = method1(data, low+1, i-1)
	n.right = method1(data, i, high)
	n.size = size(n.left) + size(n.right) + 1
	return n
}

//{10, 5, 1, 7, 40, 50}
func method2(data []int, index *int, min int, max int) *node {
	if *index >= len(data) {
		return nil
	}
	var n *node
	key := data[*index]
	if min < key && key < max {
		n = newNode(key)
		*index = *index + 1
		n.left = method2(data, index, min, key)
		n.right = method2(data, index, key, max)
		n.size = size(n.left) + size(n.right) + 1
	}
	return n
}
