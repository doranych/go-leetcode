package zero_array_transformation_ii

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	sum, k := 0, 0
	diff := make([]int, n+1)

	for i := range nums {
		for sum+diff[i] < nums[i] {
			k++
			if k > len(queries) {
				return -1
			}
			l, r, val := queries[k-1][0], queries[k-1][1], queries[k-1][2]
			if r >= i {
				diff[max(l, i)] += val
				diff[r+1] -= val
			}
		}
		sum += diff[i]
	}
	return k
}
