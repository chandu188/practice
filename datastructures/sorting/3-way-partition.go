package sorting

func _3WaySort(data []int) {
	_3wayPartitionSort(data, 0, len(data)-1)
}

func _3wayPartitionSort(data []int, lo, high int) {

	if high <= lo {
		return
	}

	lt := lo
	gt := high
	i := lo + 1
	v := data[lo]

	for i <= gt {
		if data[i] < v {
			swap(data, i, lt)
			i++
			lt++
		} else if data[i] > v {
			swap(data, i, gt)
			gt--
		} else {
			i++
		}
	}

	_3wayPartitionSort(data, lo, lt-1)
	_3wayPartitionSort(data, gt+1, high)

}
