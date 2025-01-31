package making_a_large_island

var dirs = [5]int{-1, 0, 1, 0, -1}

func largestIsland(grid [][]int) int {
	n := len(grid)
	islandMap := make(map[int]int, 0)
	id := 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				islandMap[id] = dfs(grid, i, j, id)
				id++
			}
		}
	}
	if len(islandMap) == 0 {
		return 1
	}
	if len(islandMap) == n*n {
		return n * n
	}

	maxIslandSize := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				continue
			}
			seen := map[int]bool{}
			newSize := 1
			for d := 0; d < 4; d++ {
				newI, newJ := i+dirs[d], j+dirs[d+1]
				if newI < 0 || newI >= n || newJ < 0 || newJ >= n {
					continue
				}
				tId := grid[newI][newJ]
				if tId == 0 {
					continue
				}
				if tId > 1 && !seen[tId] {
					seen[tId] = true
					newSize += islandMap[tId]
				}
			}
			maxIslandSize = max(maxIslandSize, newSize)
		}
	}

	return maxIslandSize
}

func dfs(grid [][]int, i, j, id int) int {
	if grid[i][j] != 1 {
		return 0
	}
	n := len(grid)
	grid[i][j] = id
	size := 1
	for d := 0; d < 4; d++ {
		newI, newJ := i+dirs[d], j+dirs[d+1]
		if newI < 0 || newI >= n || newJ < 0 || newJ >= n {
			continue
		}
		size += dfs(grid, newI, newJ, id)
	}
	return size
}
