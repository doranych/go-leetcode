package maximum_count_of_positive_integer_and_negative_integer

func maximumCount(nums []int) int {
	cn, cp := 0, 0
	for _, v := range nums {
		if v < 0 {
			cn++
		} else if v > 0 {
			cp++
		}
	}
	return max(cn, cp)
}
