package dynamicprogramming

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestTilingProblem(t *testing.T)  {
	tests := []struct{
		input int
		res int
	}{
		{
			input : 3,
			res: 3,
		},
		{
			input : 4,
			res: 5,
		},
	}

	for _, tc := range tests {
		act := TilingProblem(tc.input)
		assert.Equal(t, tc.res, act)
	}
	
}