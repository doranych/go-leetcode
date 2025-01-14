package minimum_length_of_string_after_operations

func minimumLength(s string) int {
	counter := [26]uint8{}

	for _, r := range s {
		counter[r-'a'] += 1
		if counter[r-'a'] == 3 {
			counter[r-'a'] = 1
		}
	}

	res := 0
	for _, c := range counter {
		res += int(c)
	}
	return res
}
