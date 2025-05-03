package maximum_candies_allocated_to_k_children

func maximumCandies(candies []int, k int64) int {
	maxN := -1
	for i := range candies {
		maxN = max(maxN, candies[i])
	}

	l, r := 0, maxN
	for l < r {
		mid := (r + l + 1) / 2
		if canAllocate(candies, k, mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func canAllocate(candies []int, k int64, num int) bool {
	maxNum := 0
	for i := range candies {
		maxNum += candies[i] / num
	}
	return int64(maxNum) >= k
}
