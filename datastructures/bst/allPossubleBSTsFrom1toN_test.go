package bst

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAllPossibelBsts(t *testing.T) {
	bsts := GenerateBSTs(3)
	assert.Equal(t, 5, len(bsts))
}