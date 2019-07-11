package dynamicprogramming

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func  TestUglyNumber(t *testing.T)  {
	un := UglyNumber(10)
	assert.Equal(t,3 , un)
}