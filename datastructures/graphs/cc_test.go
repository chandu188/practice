package graphs

import "testing"
import "github.com/stretchr/testify/assert"

func TestCC(t *testing.T) {
	g := NewGraph(13)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 5)
	g.addEdge(0, 6)
	g.addEdge(5, 3)
	g.addEdge(5, 4)
	g.addEdge(4, 6)

	g.addEdge(7, 8)

	g.addEdge(9, 10)
	g.addEdge(9, 12)
	g.addEdge(9, 11)
	g.addEdge(12, 11)

	cc := NewCC(g)
	assert.Equal(t, 3, cc.Count())
	assert.Equal(t, 0, cc.Id(6))
	assert.Equal(t, 2, cc.Id(12))

}
