package minimum_cost_walk_in_weighted_graph

var (
	parents []int
	comp    []int
	depth   []int
)

func minimumCost(n int, edges [][]int, queries [][]int) []int {
	parents = make([]int, n)
	comp = make([]int, n)
	depth = make([]int, n)

	for i := range parents {
		parents[i] = -1
		comp[i] = 1<<32 - 1
	}

	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	for _, edge := range edges {
		comp[find(edge[0])] &= edge[2]
	}

	ans := make([]int, len(queries))
	for i := range queries {
		s, e := queries[i][0], queries[i][1]
		if find(s) != find(e) {
			ans[i] = -1
			continue
		}
		ans[i] = comp[find(s)]
	}
	return ans

}

func find(n int) int {
	if parents[n] == -1 {
		return n
	}
	parents[n] = find(parents[n])
	return parents[n]
}

func union(n1, n2 int) {
	r := find(n1)
	r2 := find(n2)
	if r == r2 {
		return
	}
	if depth[r] < depth[r2] {
		r, r2 = r2, r
	}
	parents[r2] = r

	if depth[r] == depth[r2] {
		depth[r]++
	}
}
