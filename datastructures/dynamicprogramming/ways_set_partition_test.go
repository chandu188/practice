package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPartition(t *testing.T) {

	res := SetPartition(1, 1)
	assert.Equal(t, 1, res)

	res = SetPartition(3, 3)
	assert.Equal(t, 5, res)

	res = SetPartition(4, 4)
	assert.Equal(t, 15, res)

	res = SetPartition(5, 5)
	assert.Equal(t, 52, res)
}
