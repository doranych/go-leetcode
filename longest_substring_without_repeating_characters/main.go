package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	curMax := 0
	curString := make([]rune, 0)
	lastStartPos := 0
	curHash := make(map[rune]int)
	for i, c := range s {
		pos, ok := curHash[c]
		if !ok || pos < lastStartPos {
			curString = append(curString, c)
		} else if pos >= lastStartPos {
			lastStartPos = pos
			newStart := len(curString) - (i - pos)
			curString = curString[newStart:]
		}

		curHash[c] = i
		curMax = max(curMax, len(curString))
	}
	return max(curMax, len(curString))
}
