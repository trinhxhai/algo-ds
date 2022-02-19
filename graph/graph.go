package graph

type Edge struct {
	src, des, w, i int
}

type Graph struct {
	// adjacency list
	adj []Edge
}

func (graph *Graph) Add(src, des int) {
	graph.adj = append(graph.adj, Edge{src: src, des: des})
}
