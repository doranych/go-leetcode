package maximum_ascending_subarray_sum

func maxAscendingSum(nums []int) int {
	maxSum := 0
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			curSum += nums[i]
		} else {
			maxSum = max(maxSum, curSum)
			curSum = nums[i]
		}
	}
	return max(maxSum, curSum)
}
