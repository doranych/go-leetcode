package count_the_number_of_fair_pairs

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	return lb(nums, upper+1) - lb(nums, lower)
}

func lb(nums []int, v int) int64 {
	l, r := 0, len(nums)-1
	res := 0

	for l < r {
		sum := nums[l] + nums[r]
		if sum < v {
			res += r - l
			l++
		} else {
			r--
		}
	}

	return int64(res)
}
