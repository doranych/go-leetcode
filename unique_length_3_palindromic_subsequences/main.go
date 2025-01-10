package unique_length_3_palindromic_subsequences

func countPalindromicSubsequence(s string) int {
	ans := 0
	indexes := make(map[rune][]int)
	for ix, val := range s {
		if indexes[val] == nil {
			indexes[val] = make([]int, 0)
		}
		indexes[val] = append(indexes[val], ix)
	}

	for _, ix := range indexes {
		l := ix[0]
		r := ix[len(ix)-1]
		for _, ix2 := range indexes {
			for _, i := range ix2 {
				if i > l && i < r {
					ans++
					break
				}
			}
		}
	}
	return ans
}
