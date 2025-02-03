package longest_strictly_increasing_or_strictly_decreasing_subarray

func longestMonotonicSubarray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	dir, maxMono, counter := 0, 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			if dir != 1 {
				maxMono = max(maxMono, counter)
				counter = 1
			}
			counter++
			dir = 1
		} else if nums[i-1] > nums[i] {
			if dir != -1 {
				maxMono = max(maxMono, counter)
				counter = 1
			}
			counter++
			dir = -1
		} else {
			maxMono = max(maxMono, counter)
			counter = 1
		}
	}
	return max(maxMono, counter)
}
