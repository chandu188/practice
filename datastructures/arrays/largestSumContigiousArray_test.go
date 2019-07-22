package arrays

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxContigiousSumSubArray(t *testing.T) {
	res := LargestSumContigiousArray([]int{-2, -3, 4, -1, -2, 1, 5, -3})
	assert.Equal(t, 7, res)
}
