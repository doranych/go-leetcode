package largest_local_values_in_a_matrix

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	output := make([][]int, n-2)
	for i := 1; i < n-1; i++ {
		output[i-1] = make([]int, n-2)
		for j := 1; j < n-1; j++ {
			maxN := 0
			for k := i - 1; k <= i+1; k++ {
				for m := j - 1; m <= j+1; m++ {
					maxN = max(grid[k][m], maxN)
				}
			}
			output[i-1][j-1] = maxN
		}
	}
	return output
}
