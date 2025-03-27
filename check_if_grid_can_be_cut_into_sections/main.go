package check_if_grid_can_be_cut_into_sections

import "sort"

func checkValidCuts(_ int, rectangles [][]int) bool {
	checkCuts := func(dim int) bool {
		gapCount := 0
		sort.Slice(rectangles, func(i, j int) bool {
			return rectangles[i][dim] <= rectangles[j][dim]
		})
		end := rectangles[0][dim+2]
		for i := 1; i < len(rectangles); i++ {
			if end <= rectangles[i][dim] {
				gapCount++
			}
			end = max(end, rectangles[i][dim+2])
		}
		return gapCount >= 2
	}
	return checkCuts(0) || checkCuts(1)
}
