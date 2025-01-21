package grid_game

import "math"

func gridGame(grid [][]int) int64 {
	sum := math.MaxInt64
	sumFirst := 0
	sumSecond := 0
	for i := 0; i < len(grid[0]); i++ {
		sumFirst += grid[0][i]
	}

	for i := range grid[0] {
		sumFirst -= grid[0][i]
		sum = min(sum, max(sumFirst, sumSecond))
		sumSecond += grid[1][i]
	}

	return int64(sum)
}
