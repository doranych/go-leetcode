package counting_words_with_a_given_prefix

func prefixCount(words []string, pref string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		if len(pref) <= len(words[i]) && pref == words[i][:len(pref)] {
			res++
		}
	}
	return res
}
