package graphs

import "github.com/chandu188/practice/datastructures/stack"

type dfs struct {
	s      int
	g      *graph
	marked []bool
	edgeTo []int
}

func NewDFS(g *graph, s int) *dfs {
	d := &dfs{
		g:      g,
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}
	d.visit(s)
	return d
}

func (d *dfs) visit(u int) {
	d.marked[u] = true
	for _, v := range d.g.adjVertices(u) {
		if !d.marked[v] {
			d.visit(v)
			d.edgeTo[v] = u
		}
	}
}

func (d *dfs) hasPathTo(v int) bool {
	return d.marked[v]
}

func (d *dfs) pathTo(v int) stack.Stack {
	if !d.hasPathTo(v) {
		return nil
	}
	stk := stack.NewStack()
	for s := v; s != d.s; s = d.edgeTo[s] {
		stk.Push(s)
	}
	stk.Push(d.s)
	return stk
}
