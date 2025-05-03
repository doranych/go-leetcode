package reverse_vowels_of_a_string

func reverseVowels(s string) string {
	str := []rune(s)
	i, j := 0, len(s)-1
	for {
		if i >= j {
			break
		}
		_, iOk := vowels[str[i]]
		_, jOk := vowels[str[j]]
		if !iOk {
			i++
		}
		if !jOk {
			j--
		}
		if iOk && jOk {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	return string(str)
}

var vowels = map[rune]bool{
	'A': true,
	'a': true,
	'E': true,
	'e': true,
	'I': true,
	'i': true,
	'O': true,
	'o': true,
	'U': true,
	'u': true,
}
