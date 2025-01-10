package shifting_letters_ii

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	shiftCharacters := make([]int, len(s))
	str := []byte(s)
	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			shiftCharacters[start] = shiftCharacters[start] + 1
			if end+1 < n {
				shiftCharacters[end+1] = shiftCharacters[end+1] - 1
			}
		} else {
			shiftCharacters[start] = shiftCharacters[start] - 1
			if end+1 < n {
				shiftCharacters[end+1] = shiftCharacters[end+1] + 1
			}
		}
	}
	num := 0
	for i, shift := range shiftCharacters {
		num += shift
		str[i] += uint8(num % 26)
		if str[i] > 'z' {
			str[i] = (str[i] - 'z' - 1) + 'a'
		}
		if str[i] < 'a' {
			str[i] = 'z' - ('a' - str[i] - 1)
		}
	}
	return string(str)
}
