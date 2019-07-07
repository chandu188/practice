package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test2Stack1Array(t *testing.T) {
	ts := New2Stacks(5)
	ts.Push1(5)
	ts.Push2(10)
	ts.Push2(15)
	ts.Push1(11)
	ts.Push2(7)
	assert.Equal(t, 11, ts.Pop1())
	ts.Push2(40)
	assert.Equal(t, 40, ts.Pop2())
}
