package graphs

import "testing"
import "github.com/stretchr/testify/assert"

func TestGraph(t *testing.T) {
	g := NewGraph(4)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 3)
	assert.Equal(t, 2, len(g.adjVertices(1)))
	assert.Equal(t, 2, len(g.adjVertices(0)))
	assert.Equal(t, 1, len(g.adjVertices(2)))
	assert.Equal(t, 1, len(g.adjVertices(3)))
	assert.Equal(t, 3, g.E())
	assert.Equal(t, 4, g.V())
}
