package graphs

import "testing"
import "github.com/stretchr/testify/assert"

func TestDFS(t *testing.T) {
	g := NewGraph(6)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 3)
	g.addEdge(4, 5)

	dfs := NewDFS(g, 0)
	assert.Equal(t, true, dfs.hasPathTo(1))
	assert.Equal(t, false, dfs.hasPathTo(4))
	stk := dfs.pathTo(3)
	res := []int{0, 1, 3}
	assert.Equal(t, len(res), stk.Size())
	for _, v := range res {
		assert.Equal(t, v, stk.Pop())
	}
}
