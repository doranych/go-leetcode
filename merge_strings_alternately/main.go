package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0
	res := make([]byte, 0, len(word1)+len(word2))
	for {
		if i < len(word1) {
			res = append(res, word1[i])
			i++
		}
		if j < len(word2) {
			res = append(res, word2[j])
			j++
		}
		if i >= len(word1) && j >= len(word2) {
			break
		}
	}
	return string(res)
}
