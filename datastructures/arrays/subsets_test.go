package arrays

import "testing"
import "github.com/stretchr/testify/assert"

func TestSubSets(t *testing.T) {
	in := []int{1, 2, 3}
	ss := Subsets(in)
	assert.Equal(t, 8, len(ss))
}
