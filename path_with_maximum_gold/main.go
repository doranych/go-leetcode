package path_with_maximum_gold

func getMaximumGold(grid [][]int) int {
	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			runner(grid, i, j, grid[i][j], &result)
		}
	}

	return result
}

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func runner(grid [][]int, x, y, curr int, result *int) {
	*result = max(*result, curr)

	tmp := grid[x][y]
	grid[x][y] = 0

	for _, v := range dirs {
		xx, yy := x+v[0], y+v[1]
		if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[0]) || grid[xx][yy] == 0 {
			continue
		}
		runner(grid, xx, yy, curr+grid[xx][yy], result)
	}

	grid[x][y] = tmp
}
