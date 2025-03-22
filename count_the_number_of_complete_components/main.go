package count_the_number_of_complete_components

func countCompleteComponents(n int, edges [][]int) int {
	g := make([][]int, n)
	for i := range n {
		g[i] = make([]int, 0)
	}
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}
	visited := make([]bool, n)
	res := 0

	for v := range n {
		if visited[v] {
			continue
		}
		comp := make([]int, 0)
		q := []int{v}
		visited[v] = true

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			comp = append(comp, cur)
			for _, neighbour := range g[cur] {
				if visited[neighbour] {
					continue
				}
				q = append(q, neighbour)
				visited[neighbour] = true
			}
		}
		isComp := true
		for _, node := range comp {
			if len(g[node]) != len(comp)-1 {
				isComp = false
				break
			}
		}
		if isComp {
			res++
		}
	}

	return res
}
