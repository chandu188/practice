package dynamicprogramming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUglyNumber(t *testing.T) {
	un := UglyNumber(10)
	assert.Equal(t, 3, un)
}
