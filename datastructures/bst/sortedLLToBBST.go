package bst

type llNode struct {
	data int
	next *llNode
}

type LinkedList struct {
	front *llNode
	rear  *llNode
	size  int
}

func newLLNode(data int) *llNode {
	return &llNode{data: data}
}

func (ll *LinkedList) Append(data int) {
	if ll.front == nil {
		ll.front = newLLNode(data)
		ll.rear = ll.front
		ll.size = ll.size + 1
		return
	}
	n := newLLNode(data)
	ll.rear.next = n
	ll.rear = n
	ll.size = ll.size + 1
}

func newLLFrom(data []int) *LinkedList {
	ll := &LinkedList{}
	for _, ele := range data {
		ll.Append(ele)
	}
	return ll
}

func SortedLLTOBBST(ll *LinkedList) *Bst {
	return &Bst{
		root: sortedLLTOBBST(&ll.front, ll.size),
	}
}

/*

Input:  Linked List 1->2->3
Output: A Balanced BST
     2
   /  \
  1    3


Input: Linked List 1->2->3->4->5->6->7
Output: A Balanced BST
        4
      /   \
     2     6
   /  \   / \
  1   3  5   7
*/
func sortedLLTOBBST(head **llNode, n int) *node {
	if n <= 0 {
		return nil
	}

	left := sortedLLTOBBST(head, n/2)

	node := newNode((*head).data)
	*head = (*head).next
	node.left = left

	node.right = sortedLLTOBBST(head, n-n/2-1)
	return node
}
