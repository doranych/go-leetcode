package minimum_operations_to_make_a_uni_value_grid

import "sort"

func minOperations(grid [][]int, x int) int {
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	a := make([]int, 0, len(grid)*len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			a = append(a, grid[i][j])
		}
	}
	sort.Ints(a)
	med := a[len(a)/2]
	res := 0
	for _, n := range a {
		if n%x != med%x {
			return -1
		}
		res += abs(med-n) / x
	}

	return res
}
