package check_if_array_is_sorted_and_rotated

func check(nums []int) bool {
	inv := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			inv++
		}
	}
	if nums[0] < nums[len(nums)-1] {
		inv++
	}

	return inv <= 1
}
