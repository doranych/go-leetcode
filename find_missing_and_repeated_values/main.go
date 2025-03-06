package find_missing_and_repeated_values

/**
optimal solution would be some sort of math trick, which i don't wanna think about
*/

func findMissingAndRepeatedValues(grid [][]int) []int {
	ans := make([]int, 2)
	m := make([]uint, len(grid)*len(grid))
	for i := range grid {
		for j := range grid[i] {
			m[grid[i][j]-1]++
		}
	}
	for i := range m {
		if m[i] == 2 {
			ans[0] = i + 1
		}
		if m[i] == 0 {
			ans[1] = i + 1
		}
		if ans[0] != 0 && ans[1] != 0 {
			return ans
		}
	}
	return ans
}
