package sorting

func SelectionSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		mi := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[mi] {
				mi = j
			}
		}
		swap(a, i, mi)
	}
}

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
