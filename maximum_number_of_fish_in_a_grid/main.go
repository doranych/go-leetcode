package maximum_number_of_fish_in_a_grid

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	maxCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 && !visited[i][j] {
				maxCount = max(maxCount, dfs(grid, visited, i, j))
			}
		}
	}
	return maxCount
}

func dfs(grid [][]int, visited [][]bool, i, j int) int {
	m, n := len(grid), len(grid[0])
	q := [][2]int{{i, j}}
	total := 0
	visited[i][j] = true
	dirs := [5]int{-1, 0, 1, 0, -1}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		total += grid[c[0]][c[1]]
		for d := 0; d < 4; d++ {
			newX, newY := c[0]+dirs[d], c[1]+dirs[d+1]
			if newX < 0 || newX >= m || newY < 0 || newY >= n {
				continue
			}
			if grid[newX][newY] == 0 || visited[newX][newY] {
				continue
			}
			q = append(q, [2]int{newX, newY})
			visited[newX][newY] = true
		}
	}
	return total
}
