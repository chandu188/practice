package stack

func ReverseStack(st Stack) {
	reverseStack(st)
}

func insertAtBottom(st Stack, n int) {
	if st.IsEmpty() {
		st.Push(n)
		return
	}

	ele := st.Pop()
	insertAtBottom(st, n)
	st.Push(ele)
}

func reverseStack(st Stack) {
	if st.IsEmpty() {
		return
	}
	ele := st.Pop().(int)
	reverseStack(st)
	println(ele)
	insertAtBottom(st, ele)
}
