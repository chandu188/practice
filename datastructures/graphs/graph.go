package graphs

type graph struct {
	noOfVertices int
	adj          [][]int
	noOfedges    int
}

func NewGraph(v int) *graph {
	var adj [][]int
	for i := 0; i < v; i++ {
		adj = append(adj, make([]int, 0))
	}
	return &graph{
		noOfVertices: v,
		adj:          adj,
	}
}

func (g *graph) V() int {
	return g.noOfVertices
}

func (g *graph) E() int {
	return g.noOfedges
}

func (g *graph) addEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
	g.noOfedges++
}

func (g *graph) adjVertices(u int) []int {
	return g.adj[u]
}
