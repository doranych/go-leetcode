package the_k_th_lexicographical_string_of_all_happy_strings_of_length_n

import "math"

var letters = []string{"a", "b", "c"}

func getHappyString(n int, k int) string {
	c := int(3 * math.Pow(2, float64(n-1)))
	if k > c {
		return ""
	}
	resIndex := 0
	q := []string{""}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		if len(s) == n {
			resIndex++
			if resIndex == k {
				return s
			}
			continue
		}
		for _, l := range letters {
			if len(s) > 0 && s[len(s)-1] == l[0] {
				continue
			}
			newS := s + string(l)
			q = append(q, newS)
		}
	}
	return ""
}
