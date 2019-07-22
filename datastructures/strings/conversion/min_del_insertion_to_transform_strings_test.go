package conversion

import "testing"
import "github.com/stretchr/testify/assert"

func TestMinDelInsertionsToTransfrom(t *testing.T) {
	tcs := []struct {
		s1  string
		s2  string
		res int
	}{
		{
			s1:  "geeksforgeeks",
			s2:  "geeks",
			res: 8,
		},
		{
			s1:  "heap",
			s2:  "pea",
			res: 3,
		},
	}

	for _, tc := range tcs {
		res := MinDelInsertionsToTransfrom(tc.s1, tc.s2)
		assert.Equal(t, tc.res, res)
	}

}
