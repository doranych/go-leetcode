package redundant_connection

import "slices"

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	// we have 1 more edge than nodes
	visited := make([]bool, n)
	parents := make([]int, n)
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[edges[i][0]-1] = append(g[edges[i][0]-1], edges[i][1]-1)
		g[edges[i][1]-1] = append(g[edges[i][1]-1], edges[i][0]-1)
		parents[i] = -1
	}

	cycleStart := -1
	dfs(g, visited, parents, &cycleStart, 0)

	cycle := make([]int, 0)
	node := cycleStart
	for {
		cycle = append(cycle, node)
		node = parents[node]
		if node == cycleStart {
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if slices.Contains(cycle, edges[i][0]-1) && slices.Contains(cycle, edges[i][1]-1) {
			return edges[i]
		}
	}
	// should never happen
	return []int{}
}

func dfs(g [][]int, visited []bool, parents []int, cycleStart *int, cur int) {
	visited[cur] = true
	for _, v := range g[cur] {
		if !visited[v] {
			parents[v] = cur
			dfs(g, visited, parents, cycleStart, v)
		} else if v != parents[cur] && *cycleStart == -1 {
			*cycleStart = v
			parents[v] = cur
		}
	}
}
