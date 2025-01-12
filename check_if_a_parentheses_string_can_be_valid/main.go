package check_if_a_parentheses_string_can_be_valid

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}
	bLeft := 0
	bRight := 0

	for i, r := range s {
		if r == '(' || locked[i] == '0' {
			bLeft++
		} else {
			bLeft--
		}

		if bLeft < 0 {
			return false
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			bRight++
		} else {
			bRight--
		}

		if bRight < 0 {
			return false
		}
	}

	return true
}
