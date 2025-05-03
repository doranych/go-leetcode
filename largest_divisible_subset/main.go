package largest_divisible_subset

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)
	memo := make(map[int][]int, 0)
	var eds func(int) []int
	eds = func(idx int) []int {
		if v, ok := memo[idx]; ok {
			return v
		}
		tail := nums[idx]
		maxS := []int{}
		for i := range idx {
			if tail%nums[i] == 0 {
				ss := eds(i)
				if len(maxS) < len(ss) {
					maxS = ss
				}
			}
		}

		maxS = append([]int{}, maxS...)
		maxS = append(maxS, tail)
		memo[idx] = maxS
		return maxS
	}
	res := []int{}
	for i := range len(nums) {
		ss := eds(i)
		if len(res) <= len(ss) {
			res = ss
		}
	}
	return res
}
