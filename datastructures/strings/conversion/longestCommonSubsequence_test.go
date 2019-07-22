package conversion

import "testing"
import "github.com/stretchr/testify/assert"

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		s1  string
		s2  string
		res int
	}{
		{
			s1:  "",
			s2:  "",
			res: 0,
		},
		{
			s1:  "ABC",
			s2:  "AAAABCDEF",
			res: 3,
		},
		{
			s1:  "ABCDGH",
			s2:  "AEDFHR",
			res: 3,
		}, {
			s1:  "AGGTAB",
			s2:  "GXTXAYB",
			res: 4,
		},
	}

	for _, tc := range testCases {
		res := longestCommonSubsequence(tc.s1, tc.s2)
		assert.Equal(t, res, tc.res)
	}

}
