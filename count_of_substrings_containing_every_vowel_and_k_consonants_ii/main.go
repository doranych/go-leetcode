package count_of_substrings_containing_every_vowel_and_k_consonants_ii

import "slices"

var vow = []byte{'a', 'e', 'i', 'o', 'u'}

func countOfSubstrings(word string, k int) int64 {
	return atLeast(word, k) - atLeast(word, k+1)
}
func atLeast(word string, k int) int64 {
	res := 0
	start, end := 0, 0
	vc := make(map[byte]int, 5)
	cc := 0
	for end < len(word) {
		if slices.Contains(vow, word[end]) {
			vc[word[end]]++
		} else {
			cc++
		}
		for len(vc) == 5 && cc >= k {
			res += len(word) - end
			if slices.Contains(vow, word[start]) {
				vc[word[start]]--
				if v, ok := vc[word[start]]; ok && v == 0 {
					delete(vc, word[start])
				}
			} else {
				cc--
			}
			start++
		}
		end++
	}
	return int64(res)
}
