package count_prefix_and_suffix_pairs_i

func countPrefixSuffixPairs(words []string) int {
	n := len(words)
	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if len(words[j]) >= len(words[i]) {
				if words[i] == words[j][:len(words[i])] && words[i] == words[j][len(words[j])-len(words[i]):] {
					res++
				}
			}
		}
	}
	return res
}
