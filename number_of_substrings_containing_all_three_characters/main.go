package number_of_substrings_containing_all_three_characters

func numberOfSubstrings(s string) int {
	c := [3]int{}
	res := 0
	st, end := 0, 0
	for end < len(s) {
		c[s[end]-'a']++
		for isFull(c) {
			res += len(s) - end
			c[s[st]-'a']--
			st++
		}
		end++
	}
	return res
}

func isFull(c [3]int) bool {
	return c[0] > 0 && c[1] > 0 && c[2] > 0
}
