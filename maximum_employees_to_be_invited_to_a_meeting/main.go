package maximum_employees_to_be_invited_to_a_meeting

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	graph := make([][]int, n)

	for i := range favorite {
		if graph[favorite[i]] == nil {
			graph[favorite[i]] = make([]int, 0)
		}
		graph[favorite[i]] = append(graph[favorite[i]], i)
	}

	maxCycle, pairs := 0, 0
	visited := make([]bool, n)
	for i := range favorite {
		if visited[i] {
			continue
		}
		cycle := 0
		cur := i
		cMap := map[int]int{}
		for {
			if visited[cur] {
				break
			}
			visited[cur] = true
			cMap[cur] = cycle
			cycle++
			next := favorite[cur]
			if _, ok := cMap[next]; ok {
				l := cycle - cMap[next]
				maxCycle = max(l, maxCycle)
				if l == 2 {
					visNodes1 := map[int]struct{}{cur: struct{}{}, next: struct{}{}}
					visNodes2 := map[int]struct{}{cur: struct{}{}, next: struct{}{}}
					pairs += 2 + bfs(cur, visNodes1, graph) + bfs(next, visNodes2, graph)
				}
				break
			}
			cur = next
		}
	}
	return max(maxCycle, pairs)
}

func bfs(startNode int, visited map[int]struct{}, graph [][]int) int {
	q := [][2]int{{startNode, 0}}
	maxD := 0
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		i, d := c[0], c[1]
		for _, v := range graph[i] {
			if _, ok := visited[v]; ok {
				continue
			}
			visited[v] = struct{}{}
			q = append(q, [2]int{v, d + 1})
			maxD = max(maxD, d+1)
		}
	}
	return maxD
}
