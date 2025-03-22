package longest_nice_subarray

func longestNiceSubarray(nums []int) int {
	bits, start, maxLen := 0, 0, 0
	for end := range len(nums) {
		for bits&nums[end] != 0 {
			bits ^= nums[start]
			start += 1
		}
		bits |= nums[end]
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}
