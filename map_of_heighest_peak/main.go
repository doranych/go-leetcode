package map_of_heighest_peak

func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])
	heights := make([][]int, m)

	q := make([][2]int, 0, m*n)
	for i := 0; i < m; i++ {
		heights[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 0 {
				heights[i][j] = -1
			} else {
				q = append(q, [2]int{i, j})
				heights[i][j] = 0
			}
		}
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(q) > 0 {
		pos := q[0]
		q = q[1:]
		i, j := pos[0], pos[1]
		for k := 0; k < 4; k++ {
			newX, newY := i+dirs[k][0], j+dirs[k][1]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || heights[newX][newY] != -1 {
				continue
			}
			heights[newX][newY] = heights[i][j] + 1
			q = append(q, [2]int{newX, newY})
		}
	}
	return heights
}
