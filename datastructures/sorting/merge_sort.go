package sorting

func MergeSort(a []int) {
	aux := make([]int, len(a))
	mergeSort(a, aux, 0, len(a)-1)
}

func mergeSort(a []int, aux []int, lo, high int) {
	if high <= lo {
		return
	}
	mid := lo + (high-lo)/2
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid+1, high)
	merge(a, aux, lo, mid, high)
}

func merge(a []int, aux []int, low, mid, high int) {

	for k := low; k <= high; k++ {
		aux[k] = a[k]
	}

	i := low
	j := mid + 1

	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > high {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = a[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func bottomUpMergeSort(a []int) {
	n := len(a)
	aux := make([]int, len(a))

	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, n-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
