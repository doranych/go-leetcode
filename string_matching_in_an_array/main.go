package string_matching_in_an_array

import "strings"

func stringMatching(words []string) []string {
	res := make([]string, 0)
	for _, w := range words {
		for _, m := range words {
			if len(w) < len(m) {
				if strings.Contains(m, w) {
					res = append(res, w)
					break
				}
			}
		}
	}
	return res
}
