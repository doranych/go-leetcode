package count_the_hidden_sequences

func numberOfArrays(differences []int, lower int, upper int) int {
	var minH, maxH, cur int
	for _, diff := range differences {
		cur += diff
		minH = min(minH, cur)
		maxH = max(maxH, cur)
		if maxH-minH > upper-lower {
			return 0
		}
	}
	return (upper - lower) - (maxH - minH) + 1
}
