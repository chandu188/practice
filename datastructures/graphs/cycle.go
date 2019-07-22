package graphs

import "github.com/chandu188/practice/datastructures/stack"

type Cycle struct {
	marked []bool
	edgeTo []int
	cycle  stack.Stack
}

func NewCycle(g *graph) *Cycle {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())

	c := &Cycle{
		marked: marked,
		edgeTo: edgeTo,
	}

	for i := 0; i < g.V(); i++ {
		if !c.marked[i] {

			c.dfs(g, -1, i)
		}
	}
	return c
}

func (c *Cycle) dfs(g *graph, u, v int) {
	c.marked[v] = true
	for _, w := range g.adjVertices(v) {
		if c.cycle != nil {
			return
		}
		if !(c.marked[w]) {
			c.edgeTo[w] = v
			c.dfs(g, v, w)
		} else if w != u {
			c.cycle = stack.NewStack()
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Push(x)
			}
			c.cycle.Push(w)
			c.cycle.Push(v)

		}
	}
}

func (c *Cycle) hasCycle() bool {
	return c.cycle != nil
}

func (c *Cycle) getCycle() stack.Stack {
	if !c.hasCycle() {
		return nil
	}

	return c.cycle
}
