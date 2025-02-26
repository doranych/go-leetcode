package maximum_absolute_sum_of_any_subarray

import "math"

func maxAbsoluteSum(nums []int) int {
	maxSum := math.MinInt
	minSum := math.MaxInt

	curMaxSum := 0
	curMinSum := 0
	for _, n := range nums {
		curMaxSum = max(n, curMaxSum+n)
		curMinSum = min(n, curMinSum+n)
		maxSum = max(curMaxSum, maxSum)
		minSum = min(curMinSum, minSum)
	}

	maxSum = max(maxSum, int(math.Abs(float64(minSum))))
	return maxSum
}
