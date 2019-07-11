package stack

func SortStack(st Stack) {
	sortStack(st)
}

func sortedInsert(st Stack, v int) {
	if st.IsEmpty() || v > st.Top().(int) {
		st.Push(v)
		return
	}

	ele := st.Pop()
	sortedInsert(st, v)
	st.Push(ele)
}

func sortStack(st Stack) {
	if st.IsEmpty() {
		return
	}

	ele := st.Pop().(int)
	sortStack(st)
	sortedInsert(st, ele)

}
