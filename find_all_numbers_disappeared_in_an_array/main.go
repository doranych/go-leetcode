package find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n = -n
		}
		if nums[n-1] > 0 {
			nums[n-1] = -nums[n-1]
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
