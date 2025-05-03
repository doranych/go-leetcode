package reverse_words_in_a_string

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	i, j := 0, len(words)-1
	for i < j {
		if len(words[i]) == 0 {
			i++
			continue
		}
		if len(words[j]) == 0 {
			j--
			continue
		}
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	var sb strings.Builder
	for i, w := range words {
		if len(w) == 0 {
			continue
		}
		if i != len(words)-1 {
			w += " "
		}
		sb.WriteString(w)
	}
	return strings.Trim(sb.String(), " ")
}
