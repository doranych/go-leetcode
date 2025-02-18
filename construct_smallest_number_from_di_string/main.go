package construct_smallest_number_from_di_string

func smallestNumber(pattern string) string {
	res := ""
	s := make([]int, 0, 9)
	for i := range len(pattern) + 1 {
		s = append(s, i+1)
		if i == len(pattern) || pattern[i] == 'I' {
			for len(s) > 0 {
				res += string(rune('0' + s[len(s)-1]))
				s = s[:len(s)-1]
			}
		}
	}
	return res
}
