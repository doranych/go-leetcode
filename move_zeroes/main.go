package move_zeroes

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i, j = 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
