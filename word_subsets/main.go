package word_subsets

func wordSubsets(words1 []string, words2 []string) []string {
	maxCounter := make(map[rune]uint, 26)

	for _, w := range words2 {
		m := make(map[rune]uint, len(w))
		for _, r := range w {
			m[r] += 1
		}
		for b, v := range m {
			maxCounter[b] = max(v, maxCounter[b])
		}
	}
	res := make([]string, 0, len(words1))
	for _, w := range words1 {
		m := make(map[rune]uint, len(w))
		for _, r := range w {
			m[r] += 1
		}
		ret := true
		for r, v := range maxCounter {
			if m[r] < v {
				ret = false
				break
			}
		}
		if ret {
			res = append(res, w)
		}
	}
	return res
}
