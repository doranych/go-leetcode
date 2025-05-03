package count_the_number_of_good_subarrays

func countGood(nums []int, k int) int64 {
	var result, l, pairCount int64
	counts := make(map[int]int)
	for r := range nums {
		pairCount += int64(counts[nums[r]])
		counts[nums[r]]++
		for pairCount >= int64(k) {
			result += int64(len(nums) - r)
			counts[nums[l]]--
			pairCount -= int64(counts[nums[l]])
			l++
		}
	}
	return result
}
