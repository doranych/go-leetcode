package number_of_ways_to_split_array

func waysToSplitArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	curr := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		curr += nums[i]
		if curr >= sum-curr {
			count++
		}
	}
	return count
}
