package find_eventual_safe_states

func eventualSafeNodes(graph [][]int) []int {
	safeNodes := make([]int, 0)
	visited := make([]uint8, len(graph))
	for i, _ := range graph {
		if isSafe(graph, visited, i) {
			safeNodes = append(safeNodes, i)
		}
	}
	return safeNodes
}

func isSafe(graph [][]int, visited []uint8, node int) bool {
	if len(graph[node]) == 0 {
		return true
	}

	if visited[node] != 0 {
		return visited[node] == 2
	}

	visited[node]++
	for _, v := range graph[node] {
		if !isSafe(graph, visited, v) {
			return false
		}
	}
	visited[node]++
	return true
}
