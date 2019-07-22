package sorting

func QuickSort(a []int) {
	qucikSort(a, 0, len(a)-1)
}

func partition(a []int, lo, high int) int {
	i := lo + 1
	j := high

	for {
		for a[i] < a[lo] {
			i++
			if i == high {
				break
			}
		}

		for a[lo] < a[j] {
			j--
			if j == lo {
				break
			}
		}
		if j <= i {
			break
		}
		swap(a, i, j)
	}
	swap(a, lo, j)
	return j
}

func qucikSort(a []int, lo, high int) {
	if high <= lo {
		return
	}
	p := partition(a, lo, high)
	qucikSort(a, lo, p-1)
	qucikSort(a, p+1, high)
}
