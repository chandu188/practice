package arrays

/*
	[1][2][3,4,5]
	[1][2,3][4,5]
	[1][2,3,4][5]
	[1,2][3][4,5]
	[1,2][3,4][5]
	[1,2,3][4][5]
*/

func GetMinSum(data []int, k int) int {
	segs := possibleSegments(data, k, make([][]int, k), make([][][]int, 0))
	max := -1
	for _, r := range segs {
		sum := getMinSum(r)
		if sum > max {
			max = sum
		}
	}
	return max
}

func possibleSegments(data []int, k int, segs [][]int, res [][][]int) [][][]int {
	if len(data) < k {
		return res
	}

	if k == 1 {
		segs[k-1] = data
		res = append(res, append(segs[:0:0], segs...))
		return res
	}

	for i := 1; i < len(data); i++ {
		segs[k-1] = data[:i]
		res = possibleSegments(data[i:], k-1, segs, res)
	}

	return res
}

func getMinSum(data [][]int) int {
	sum := 0
	for _, row := range data {
		min := 10000000
		for _, v := range row {
			if v < min {
				min = v
			}
		}
		sum += min
	}
	return sum
}

func GetMaxMinSumKSegments(data []int, k int) int {
	dp := make([][]int, 10)
	for i := range dp {
		dp[i] = make([]int, 10)
	}
	return maxMinSumKSegments(data, 0, k, dp)

}

func maxMinSumKSegments(data []int, ind int, k int, dp [][]int) int {
	if k == 0 {
		if ind == len(data) {
			return 0
		}

		return -100000000
	}

	if ind == len(data) {
		return -100000000
	} else if dp[ind][k] != 0 {
		return dp[ind][k]
	} else {
		res := 0
		min := data[ind]
		for i := ind; i < len(data); i++ {
			min = mini(min, data[i])
			res = max(res, maxMinSumKSegments(data, i+1, k-1, dp)+min)
		}
		dp[ind][k] = res
	}
	return dp[ind][k]

}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
