package graph

type DSU struct {
	Size         int
	Par          []int
	Sz           []int
	Is_connected bool
}

func NewDSU(n int) DSU {
	dsu := DSU{
		Size: n,
		Par:  make([]int, n+1),
		Sz:   make([]int, n+1),
	}
	for i := 0; i <= n; i++ {
		dsu.Par[i] = i
		dsu.Sz[i] = 1
	}
	if n == 1 {
		dsu.Is_connected = true
	}
}

func (dsu *DSU) Find(u int) int {
	if dsu.Par[u] == u {
		return u
	}
	dsu.Par[u] = dsu.Find(dsu.Par[u])
	return dsu.Par[u]
}
func (dsu *DSU) Merge(a, b int) int {
	// return root of two after merge
	ra := dsu.Find(a)
	rb := dsu.Find(b)
	if ra != rb {
		if dsu.Sz[ra] != dsu.Sz[rb] {
			ra, rb = rb, ra
		}
		dsu.Sz[ra] += dsu.Sz[rb]
		dsu.Par[rb] = ra
		if dsu.Sz[ra] == b {
			dsu.Is_connected = true
		}
	}
	return ra
}
