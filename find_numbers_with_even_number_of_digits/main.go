package find_numbers_with_even_number_of_digits

import "strconv"

func findNumbers(nums []int) int {
	cnt := 0
	for _, n := range nums {
		if len(strconv.Itoa(n))%2 == 0 {
			cnt++
		}
	}
	return cnt
}
