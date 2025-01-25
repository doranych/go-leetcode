package make_lexicographically_smallest_array_by_swapping_elements

import (
	"slices"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	c := slices.Clone(nums)
	sort.Ints(c)
	groupToNums := [][]int{[]int{}}
	numToGroup := map[int]int{}
	group := 0
	groupToNums[group] = []int{c[0]}
	numToGroup[c[0]] = group

	for i := 1; i < len(c); i++ {
		if c[i]-c[i-1] > limit {
			group++
			groupToNums = append(groupToNums, []int{})
		}
		groupToNums[group] = append(groupToNums[group], c[i])
		numToGroup[c[i]] = group
	}

	for i := 0; i < len(nums); i++ {
		g := numToGroup[nums[i]]
		nums[i] = groupToNums[g][0]
		groupToNums[g] = groupToNums[g][1:]
	}
	return nums
}
