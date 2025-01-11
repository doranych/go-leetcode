package construct_k_palindrome_strings

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	} else if len(s) == k {
		return true
	}

	m := make(map[rune]int, 26)
	for _, r := range s {
		m[r] += 1
	}

	oddCnt := 0
	for _, v := range m {
		if v%2 == 1 {
			oddCnt++
		}
	}

	return oddCnt <= k
}
