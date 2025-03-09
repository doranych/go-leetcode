package alternating_groups_ii

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	res := 0
	alt := 1
	lastCol := colors[0]
	for i := 1; i < n+k-1; i++ {
		if colors[i%n] == lastCol {
			alt = 1
			lastCol = colors[i%n]
			continue
		}
		alt++
		if alt >= k {
			res++
		}
		lastCol = colors[i%n]
	}
	return res
}
