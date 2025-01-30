package divide_nodes_into_the_maximum_number_of_groups

func magnificentSets(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, v := range edges {
		g[v[0]-1] = append(g[v[0]-1], v[1]-1)
		g[v[1]-1] = append(g[v[1]-1], v[0]-1)
	}
	colours := make([]int, n)
	for i := 0; i < n; i++ {
		colours[i] = -1
	}
	for node := range g {
		if colours[node] != -1 {
			continue
		}
		colours[node] = 0
		if !isBipartite(g, colours, node) {
			return -1
		}
	}

	distances := make([]int, n)

	for node := range g {
		distances[node] = longestShortestPath(g, node, n)
	}

	maxGroups := 0
	visited := make([]bool, n)
	for node := range g {
		if visited[node] {
			continue
		}
		maxGroups += getNumberOfGroups(g, distances, visited, node)
	}
	return maxGroups
}

func getNumberOfGroups(g [][]int, distances []int, visited []bool, node int) int {
	maxNumber := distances[node]
	visited[node] = true
	for _, v := range g[node] {
		if visited[v] {
			continue
		}
		maxNumber = max(maxNumber, getNumberOfGroups(g, distances, visited, v))
	}
	return maxNumber
}

func longestShortestPath(g [][]int, src int, n int) int {
	distance := 0
	visited := make([]bool, n)
	visited[src] = true
	q := []int{src}
	for len(q) > 0 {
		nodesInLayers := len(q)
		for i := 0; i < nodesInLayers; i++ {
			curNode := q[0]
			q = q[1:]
			for _, v := range g[curNode] {
				if visited[v] {
					continue
				}
				visited[v] = true
				q = append(q, v)
			}
		}
		distance++
	}
	return distance
}

func isBipartite(g [][]int, colours []int, node int) bool {
	for _, v := range g[node] {
		if colours[v] == colours[node] {
			return false
		}
		if colours[v] != -1 {
			continue
		}
		colours[v] = (colours[node] + 1) % 2
		if !isBipartite(g, colours, v) {
			return false
		}
	}
	return true
}
