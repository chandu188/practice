package arrays

import "testing"
import "github.com/stretchr/testify/assert"

func TestSegments(t *testing.T) {
	in := []int{5, 7, 4, 2, 8, 1, 6}
	k := 3
	segs := make([][]int, k)
	res := make([][][]int, 0)
	res = possibleSegments(in, k, segs, res)
	assert.Equal(t, 15, len(res))
}

//GetMinSum

func TestSegmentsMaxMinSum(t *testing.T) {
	in := []int{5, 7, 4, 2, 8, 1, 6}
	k := 3
	res := GetMinSum(in, k)
	assert.Equal(t, 13, res)

	in = []int{6, 5, 3, 8, 9, 10, 4, 7, 10}
	k = 4
	res = GetMinSum(in, k)
	assert.Equal(t, 27, res)
}

// func TestSegments2(t *testing.T)  {
// 	in := []int{5, 7, 4, 2, 8, 1, 6}
// 	k:=1
// 	segs := make([][]int, k)
// 	res := make([][][]int, 0)
// 	res = possibleSegments(in, k, segs, res)
// 	assert.Equal(t, 5, len(res))
// }

func TestSegmentsDP(t *testing.T) {
	in := []int{5, 7, 4, 2, 8, 1, 6}
	k := 3
	res := GetMaxMinSumKSegments(in, k)
	assert.Equal(t, 13, res)

	in = []int{6, 5, 3, 8, 9, 10, 4, 7, 10}
	k = 4
	res = GetMaxMinSumKSegments(in, k)
	assert.Equal(t, 27, res)
}
