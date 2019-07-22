package graphs

type CC struct {
	g      *graph
	id     []int
	marked []bool
	count  int
}

func NewCC(g *graph) *CC {
	id := make([]int, g.V())
	marked := make([]bool, g.V())
	cc := &CC{
		g:      g,
		marked: marked,
		id:     id,
	}
	for i := 0; i < g.V(); i++ {
		if !cc.marked[i] {
			cc.dfs(i)
			cc.count++
		}

	}
	return cc

}

func (cc *CC) dfs(v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, w := range cc.g.adjVertices(v) {
		if !cc.marked[w] {
			cc.dfs(w)
		}
	}
}

func (cc *CC) Count() int {
	return cc.count
}

func (cc *CC) Id(v int) int {
	return cc.id[v]
}
