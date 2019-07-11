package dynamicprogramming

import "testing"
import "github.com/stretchr/testify/assert"

func TestGoldMine(t *testing.T) {

	mat := [][]int{{1, 3, 3},
                   {2, 1, 4},
                  {0, 6, 4}};
	res := GoldMine(mat, 3, 3 )
	assert.Equal(t, 12, res)

	mat = [][]int {{10, 33, 13, 15},
	{22, 21, 04, 1},
	{5, 0, 2, 3},
	{0, 6, 14, 2}};

	res = GoldMine(mat, 4, 4 )
	assert.Equal(t, 83, res)


}