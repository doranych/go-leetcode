package partition_equal_subset_sum

func canPartition(nums []int) bool {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	if total%2 != 0 {
		return false
	}
	ssSum := total / 2
	dp := make([]bool, ssSum+1)
	dp[0] = true
	for i := range nums {
		for j := ssSum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[ssSum]
}
