package minimum_recolors_to_get_k_consecutive_black_blocks

func minimumRecolors(s string, k int) int {
	curWhite := 0
	for i := 0; i < k; i++ {
		if s[i] == 'W' {
			curWhite++
		}
	}
	minWhite := curWhite
	for i := k; i < len(s); i++ {
		if s[i] == 'B' && s[i-k] == 'W' {
			curWhite--
		}
		if s[i] == 'W' && s[i-k] == 'B' {
			curWhite++
		}
		minWhite = min(curWhite, minWhite)
		if minWhite == 0 {
			return 0
		}
	}
	return minWhite
}
