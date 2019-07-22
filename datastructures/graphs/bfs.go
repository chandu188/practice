package graphs

import "github.com/chandu188/practice/datastructures/queue"

import "github.com/chandu188/practice/datastructures/stack"

type bfs struct {
	s      int
	g      *graph
	marked []bool
	edgeTo []int
}

func NewBFS(g *graph, s int) *bfs {
	bfs := &bfs{
		s:      s,
		g:      g,
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
	}
	bfs.visit(s)
	return bfs
}

func (b *bfs) visit(u int) {
	q := queue.NewQueue()
	q.Push(u)
	for !q.IsEmpty() {
		v := q.Pop()
		b.marked[v] = true
		for _, w := range b.g.adjVertices(v) {
			if !b.marked[w] {
				b.marked[w] = true
				b.edgeTo[w] = v
				q.Push(w)
			}
		}
	}

}

func (d *bfs) hasPathTo(v int) bool {
	return d.marked[v]
}

func (d *bfs) pathTo(v int) stack.Stack {
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
