package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllPossibelBsts(t *testing.T) {
	bsts := GenerateBSTs(3)
	assert.Equal(t, 5, len(bsts))
}
