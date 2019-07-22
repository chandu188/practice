package strings

import "testing"
import "github.com/stretchr/testify/assert"

func TestSubstitue(t *testing.T) {

	//ops := make([]Operation, 3)
	ops := []Operation{
		Operation{
			startIndex: 0,
			endIndex:   1,
			substitute: "cat",
		},
		Operation{
			startIndex: 3,
			endIndex:   5,
			substitute: "chandu",
		},
		Operation{
			startIndex: 6,
			endIndex:   6,
			substitute: "sathish",
		}}
	// abcdefgh -> catcchandusathishh
	str := substitute("abcdefgh", ops)
	res := "catcchandusathishh"
	assert.Equal(t, res, str)
	str = substitute("abcdefgh", []Operation{Operation{
		startIndex: 10,
		endIndex:   15,
		substitute: "sathish",
	}})
	res = "abcdefgh"
	assert.Equal(t, res, str)

}
