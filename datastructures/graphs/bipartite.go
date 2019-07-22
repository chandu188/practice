package graphs

import "github.com/chandu188/practice/datastructures/stack"

type bipartite struct {
	marked      []bool
	color       []bool
	edgeTo      []int
	isbipartite bool
	cycle       stack.Stack
}

func NewBipartite(g *graph) *bipartite {
	marked := make([]bool, g.V())
	color := make([]bool, g.V())
	edgeTo := make([]int, g.V())

	bp := &bipartite{
		marked:      marked,
		color:       color,
		edgeTo:      edgeTo,
		isbipartite: true,
	}

	for v := 0; v < g.V(); v++ {
		if !bp.marked[v] {
			bp.dfs(g, v)
		}
	}
	return bp
}

func (b *bipartite) dfs(g *graph, u int) {
	b.marked[u] = true
	for _, v := range g.adjVertices(u) {

		if b.cycle != nil {
			return
		}
		if !b.marked[v] {
			b.color[v] = !b.color[u]
			b.edgeTo[v] = u
			b.dfs(g, v)
		} else if b.color[u] == b.color[v] {
			b.isbipartite = false
			b.cycle = stack.NewStack()
			b.cycle.Push(v)
			for x := u; x != v; x = b.edgeTo[x] {
				b.cycle.Push(x)
			}
			b.cycle.Push(v)
		}

	}
}

func (b *bipartite) IsBipartite() bool {
	return b.isbipartite
}

func (b *bipartite) GetOddLengthCycle() stack.Stack {
	return b.cycle
}
