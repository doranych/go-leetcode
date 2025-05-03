package longest_palindromic_substring

func longestPalindrome(s string) string {
	maxPal := ""
	var findPalindrome = func(first int, second int) string {
		var pal = ""
		for first >= 0 && second < len(s) {
			if s[second] == s[first] {
				if second == first {
					pal = string(s[second])
				} else {
					pal = string(s[first]) + pal + string(s[second])
				}
				first--
				second++
				continue
			}
			break
		}
		return pal
	}

	for i := 0; i < len(s); i++ {
		palOdd := findPalindrome(i, i)
		palEven := findPalindrome(i-1, i)
		if len(palOdd) > len(maxPal) {
			maxPal = palOdd
		}
		if len(palEven) > len(maxPal) {
			maxPal = palEven
		}
	}
	return maxPal
}
