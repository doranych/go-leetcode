package clear_digits

func clearDigits(s string) string {
	res := ""
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			res = res[:len(res)-1]
		} else {
			res += string(s[i])
		}
	}
	return res
}
