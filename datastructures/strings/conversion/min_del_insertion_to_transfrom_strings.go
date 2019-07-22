package conversion

func MinDelInsertionsToTransfrom(a, b string) int {
	c := longestCommonSubsequence(a, b)
	minDel := len(a) - c
	minAdd := len(b) - c
	return minAdd + minDel
}

func longestCommonSubsequence(a, b string) int {
	if a == "" || b == "" {
		return 0
	}
	if a == b {
		return len(a)
	}
	if a[len(a)-1] == b[len(b)-1] {
		return 1 + longestCommonSubsequence(a[:len(a)-1], b[:len(b)-1])
	} else {
		return max(longestCommonSubsequence(a[:len(a)-1], b),
			longestCommonSubsequence(a, b[:len(b)-1]))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
