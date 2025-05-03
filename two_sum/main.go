package two_sum

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			sum := num + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
