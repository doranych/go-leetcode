package maximize_happiness_of_selected_children

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	var sum = 0
	length := len(happiness) - 1
	for i := 0; i < k; i++ {
		curHap := happiness[length-i] - i
		if curHap > 0 {
			sum += curHap
		}
	}
	return int64(sum)
}
