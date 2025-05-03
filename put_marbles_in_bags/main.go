package put_marbles_in_bags

import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	w := make([]int, n-1)
	for i := range n - 1 {
		w[i] = weights[i] + weights[i+1]
	}
	sort.Ints(w)
	res := 0
	for i := range k - 1 {
		res += w[n-2-i] - w[i]
	}
	return int64(res)
}
