package graphs

import "testing"
import "github.com/stretchr/testify/assert"

func TestCycle(t *testing.T) {
	g := NewGraph(7)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	//g.addEdge(0, 3)
	g.addEdge(0, 6)
	g.addEdge(0, 5)

	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(2, 4)
	g.addEdge(5, 4)
	g.addEdge(6, 4)

	c := NewCycle(g)

	assert.NotNil(t, c.hasCycle())
	st := c.getCycle()
	assert.Equal(t, 3, st.Size())

}
