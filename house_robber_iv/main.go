package house_robber_iv

import "slices"

func minCapability(nums []int, k int) int {
	l, r := 1, slices.Max(nums)
	n := len(nums)

	for l < r {
		m := (l + r) / 2
		p := 0

		i := 0
		for i < n {
			if nums[i] <= m {
				p++
				i += 2
			} else {
				i++
			}
		}

		if p >= k {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
