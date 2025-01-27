package course_schedule_iv

func checkIfPrerequisite(n int, req [][]int, q [][]int) []bool {
	g := make([][]int, n)
	isReachable := make([][]bool, n)
	for _, r := range req {
		if g[r[0]] == nil {
			g[r[0]] = make([]int, 0)
		}
		g[r[0]] = append(g[r[0]], r[1])
	}
	for i := 0; i < n; i++ {
		isReachable[i] = make([]bool, n)
	}

	for i := range g {
		q := []int{i}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]

			for _, n := range g[node] {
				if !isReachable[i][n] {
					isReachable[i][n] = true
					q = append(q, n)
				}
			}
		}
	}

	res := make([]bool, len(q))
	for i, qu := range q {
		res[i] = isReachable[qu[0]][qu[1]]
	}
	return res
}
