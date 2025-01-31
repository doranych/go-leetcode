package count_vowel_strings_in_ranges

import (
	"slices"
)

func vowelStrings(words []string, queries [][]int) []int {
	vowels := []uint8{'a', 'e', 'i', 'o', 'u'}
	counter := make([]int, len(words))
	c := 0
	for i := 0; i < len(words); i++ {
		if slices.Contains(vowels, words[i][0]) && slices.Contains(vowels, words[i][len(words[i])-1]) {
			c++
		}
		counter[i] = c
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		t := 0
		if q[0] > 0 {
			t = counter[q[0]-1]
		}
		res[i] = t - counter[q[1]]
		if res[i] < 0 {
			res[i] = -res[i]
		}
	}
	return res
}
