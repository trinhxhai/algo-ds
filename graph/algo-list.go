package graph

type Edge struct {
	src, des, weight int32
}

func kruskals(gNodes int32, gFrom []int32, gTo []int32, gWeight []int32) int32 {

	m := len(gFrom)
	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		edges[i] = Edge{
			src:    gFrom[i],
			des:    gTo[i],
			weight: gWeight[i],
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	var ans int32 = 0
	dsu := NewDSU(int(gNodes))

	for i := 0; i < m; i++ {
		r_src := dsu.Find(int(edges[i].src))
		r_des := dsu.Find(int(edges[i].des))
		if r_src != r_des {
			dsu.Merge(r_src, r_des)
			ans += edges[i].weight
		}
	}

	return int32(ans)
}
