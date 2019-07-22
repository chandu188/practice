package conversion

import "testing"
import "github.com/stretchr/testify/assert"

func TestConvertOneStringToAnother(t *testing.T) {
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
			s1:  "ABD",
			s2:  "BAD",
			res: 1,
		},
		{
			s1:  "EACBD",
			s2:  "EABCD",
			res: 3,
		},
	}

	for _, tc := range testCases {
		res := ConvertOneStringToAnother(tc.s1, tc.s2)
		assert.Equal(t, tc.res, res)
	}

}
