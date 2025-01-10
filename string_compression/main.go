package string_compression

import "strconv"

func compress(chars []byte) int {
	n := len(chars)
	idx := 0
	i := 0
	for i < n {
		curChar := chars[i]
		count := 0
		for i < n && chars[i] == curChar {
			count++
			i++
		}

		chars[idx] = curChar
		idx++
		if count != 1 {
			for _, dig := range []byte(strconv.Itoa(count)) {
				chars[idx] = dig
				idx++
			}
		}
	}
	return idx
}
