package conversion

import "testing"
import "github.com/stretchr/testify/assert"

func TestStringSubSequences(t *testing.T) {

	res := segments("abc")
	assert.Equal(t, 8, len(res))
}

func TestTransfromString(t *testing.T) {
	testCases := []struct {
		s1  string
		s2  string
		res bool
	}{
		{
			s1:  "daBcd",
			s2:  "ABC",
			res: true,
		},
		{
			s1:  "argaju",
			s2:  "RAJ",
			res: true,
		},
	}

	for _, tc := range testCases {
		res := TransfromString(tc.s1, tc.s2)
		assert.Equal(t, tc.res, res)
	}
}
