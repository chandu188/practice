package sorting

func InsertionSort(a []int) {
	n := len(a)

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				swap(a, j, j-1)
			} else {
				break
			}
		}
	}
}
