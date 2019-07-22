package conversion

/*
	Given two strings A and B, the task is to convert A to B if possible.
	The only operation allowed is to put any character from A and insert it at front.
	Find if it’s possible to convert the string.
	If yes, then output minimum no. of operations required for transformation.

	A = "EACBD", B = "EABCD"

	==> BEACD, EABCD ==>

	solution:

	Below is complete algorithm.
1) Find if A can be transformed to B or not by first creating a count array for all characters of A, then checking with B if B has same count for every character.
2) Initialize result as 0.
2) Start traversing from end of both strings.
……a) If current characters of A and B match, i.e., A[i] == B[j]
………then do i = i-1 and j = j-1
……b) If current characters don’t match, then search B[j] in remaining
………A. While searching, keep incrementing result as these characters
………must be moved ahead for A to B transformation.
*/
func ConvertOneStringToAnother(a, b string) int {
	if !isPossibleToTransform(a, b) {
		return -1
	}

	i := len(a) - 1
	j := len(b) - 1
	ops := 0
	for i >= 0 {
		for i >= 0 && a[i] != b[j] {
			i--
			ops++
		}
		if i >= 0 {
			i--
			j--
		}
	}

	return ops
}

func isPossibleToTransform(a, b string) bool {
	count := make([]int, 256)
	for _, ch := range a {
		count[ch]++
	}
	for _, ch := range b {
		count[ch]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
