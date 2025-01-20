package first_completely_painted_row_or_column

import "math"

func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	intM := make(map[int]int, m*n)
	result := math.MaxInt32

	for i := 0; i < len(arr); i++ {
		intM[arr[i]] = i
	}

	for i := 0; i < m; i++ {
		lastIdx := -1
		for j := 0; j < n; j++ {
			lastIdx = max(lastIdx, intM[mat[i][j]])
		}
		result = min(result, lastIdx)
	}

	for i := 0; i < n; i++ {
		lastIdx := -1
		for j := 0; j < m; j++ {
			lastIdx = max(lastIdx, intM[mat[j][i]])
		}
		result = min(result, lastIdx)
	}

	return result
}
