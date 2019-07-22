package graphs

import "testing"
import "github.com/stretchr/testify/assert"

func TestBipartite(t *testing.T) {
	g := NewGraph(7)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 3)
	g.addEdge(0, 6)
	g.addEdge(0, 5)

	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(2, 4)
	g.addEdge(5, 4)
	g.addEdge(6, 4)

	bpp := NewBipartite(g)
	assert.Equal(t, false, bpp.IsBipartite())
	assert.NotNil(t, bpp.GetOddLengthCycle())
	cycle := bpp.GetOddLengthCycle()
	res := []int{0, 1, 3, 0}

	for _, v := range res {
		assert.Equal(t, v, cycle.Pop())
	}

}

func TestBipartite2(t *testing.T) {
	g := NewGraph(7)
	g.addEdge(0, 1)
	g.addEdge(0, 2)

	g.addEdge(0, 6)
	g.addEdge(0, 5)

	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(2, 4)
	g.addEdge(5, 4)
	g.addEdge(6, 4)

	bpp := NewBipartite(g)

	assert.Equal(t, true, bpp.IsBipartite())
	assert.Equal(t, nil, bpp.GetOddLengthCycle())
}
