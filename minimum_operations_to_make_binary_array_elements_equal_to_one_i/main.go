package minimum_operations_to_make_binary_array_elements_equal_to_one_i

func minOperations(nums []int) int {
	counter := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			nums[i] = 1
			nums[i+1] = nums[i+1] ^ 1
			nums[i+2] = nums[i+2] ^ 1
			counter++
		}
	}
	if nums[len(nums)-1] == 0 || nums[len(nums)-2] == 0 {
		return -1
	}
	return counter
}
