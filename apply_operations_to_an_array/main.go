package apply_operations_to_an_array

func applyOperations(nums []int) []int {
	n := len(nums)
	idx := 0
	for i := range n {
		if i < n-1 && nums[i] == nums[i+1] && nums[i] != 0 {
			nums[i] *= 2
			nums[i+1] = 0
		}
		if nums[i] != 0 {
			if i != idx {
				nums[i], nums[idx] = nums[idx], nums[i]
			}
			idx++
		}
	}
	return nums
}

/**
same time & memory according to leetcode submission details
twice easier to understand
func applyOperations(nums []int) []int {
    res := make([]int, 0, len(nums))
    for i := 0; i < len(nums); i++ {
        if i < len(nums)-1 && nums[i] == nums[i+1] {
            nums[i] *=2
            nums[i+1] = 0
        }
        if nums[i] != 0 {
            res = append(res, nums[i])
        }
    }
    for i := len(res); i < len(nums); i++ {
        res = append(res, 0)
    }
    return res
}
*/
