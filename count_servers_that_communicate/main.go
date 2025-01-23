package count_servers_that_communicate

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	counterR := make([]int, m)
	counterC := make([]int, n)
	c := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				counterR[i]++
				counterC[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if counterR[i] > 1 || counterC[j] > 1 {
					c++
				}
			}
		}
	}
	return c
}
