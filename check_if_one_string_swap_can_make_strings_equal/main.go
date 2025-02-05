package check_if_one_string_swap_can_make_strings_equal

import "maps"

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	c := 0
	d1 := make(map[uint8]uint8, 0)
	d2 := make(map[uint8]uint8, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			c++
			d1[s1[i]]++
			d2[s2[i]]++
		}
		if c > 2 {
			return false
		}
	}
	return (c == 0 || c == 2) && maps.Equal(d1, d2)
}
