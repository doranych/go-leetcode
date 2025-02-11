package remove_all_occurrences_of_a_substring

import "slices"

func removeOccurrences(s string, part string) string {
	l := len(part)
	res := make([]byte, 0, len(s))
	b := []byte(part)
	for i := 0; i < len(s); i++ {
		res = append(res, s[i])
		if len(res) >= l && slices.Equal(res[len(res)-l:], b) {
			res = res[:len(res)-l]
		}
	}
	return string(res)
}
