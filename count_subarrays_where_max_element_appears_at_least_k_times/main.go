package count_subarrays_where_max_element_appears_at_least_k_times

import "slices"

func countSubarrays(nums []int, k int) int64 {
	maxN := slices.Max(nums)
	res := 0
	l := 0
	maxC := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == maxN {
			maxC++
		}
		if maxC == k {
			res += len(nums) - i
		}
		for maxC >= k {
			if nums[l] == maxN {
				maxC--
			}
			if maxC >= k {
				res += len(nums) - i
			}
			l++
		}
	}

	return int64(res)
}
