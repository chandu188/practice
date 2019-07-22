package arrays

func LargestSumContigiousArray(data []int) int {
	max_ending_here := data[0]
	max_so_far := data[0]
	var start, end, s int
	for i, v := range data[1:] {
		max_ending_here = max_ending_here + v
		if max_ending_here < 0 {
			max_ending_here = 0
			s = i + 1
		}
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
			start = s
			end = i
		}
	}
	println(start, end)
	return max_so_far
}
