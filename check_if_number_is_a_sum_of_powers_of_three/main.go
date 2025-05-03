package check_if_number_is_a_sum_of_powers_of_three

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n = n / 3
	}
	return true
}
