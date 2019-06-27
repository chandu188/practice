package bst

import ("container/list"
"sort"
)
/*
Input:	  10   												  				
         /  \																
        2    7												
       / \												
      8   4												

output:    8
         /  \
        4    10
       / \
      2   7

*/


func btToBSTFrom(data []int) *Bst {
	root := arrayToBT(data)
	elems := inorder(make([]int, 0), root)
	sort.Ints(elems)
	i := 0
	btToBST(elems, &i, root)
	return &Bst{root: root}
}

func btToBST(data []int, index *int, n *node) {
	if n == nil {
		return
	}

	btToBST(data, index, n.left)
	n.data = data[*index]
	*index = *index + 1
	btToBST(data, index, n.right)
	return
}

func arrayToBT(data []int) *node {
	queue:= list.New()
	root := newNode(data[0])
	queue.PushBack(root)
	i :=1
	for queue.Len() > 0 && i < len(data){
		e := queue.Front()
		n := e.Value.(*node)
		l := newNode(data[i])
		i ++
		n.left = l
		if i >= len(data) {
			break
		}
		r := newNode(data[i])
		i++
		n.right = r 
		queue.PushBack(l)
		queue.PushBack(r)
		queue.Remove(e)
	}
	return root 
}

