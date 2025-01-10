package increasing_triplet_subsequence

import "math"

func increasingTriplet(nums []int) bool {
	firstMin, secondMin := math.MaxInt32, math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if nums[i] <= firstMin {
			firstMin = nums[i]
		} else if nums[i] <= secondMin {
			secondMin = nums[i]
		} else {
			return true
		}
	}
	return false
}
