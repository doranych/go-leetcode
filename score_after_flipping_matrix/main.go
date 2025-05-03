package score_after_flipping_matrix

import (
	"strconv"
	"strings"
)

func matrixScore(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			grid[i] = move(grid[i], n)
		}
	}
	for j := 0; j < n; j++ {
		var c1, c2 int
		for k := 0; k < m; k++ {
			if grid[k][j] == 1 {
				c1++
				continue
			}
			c2++
		}
		if c2 > c1 {
			for k := 0; k < m; k++ {
				grid[k][j] = flip(grid[k][j])
			}
		}
	}

	sum := 0
	for i := 0; i < m; i++ {
		sum += convertToInt(grid[i])
	}
	return sum
}

func move(row []int, n int) []int {
	for i := 0; i < n; i++ {
		row[i] = flip(row[i])
	}
	return row
}

func convertToInt(row []int) int {
	b := strings.Builder{}
	for i := 0; i < len(row); i++ {
		b.Write([]byte(strconv.Itoa(row[i])))
	}
	v, _ := strconv.ParseInt(b.String(), 2, 64)
	return int(v)
}

func flip(v int) int {
	if v == 1 {
		return 0
	}
	return 1
}
