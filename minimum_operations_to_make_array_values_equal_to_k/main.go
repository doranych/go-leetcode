package minimum_operations_to_make_array_values_equal_to_k

func minOperations(nums []int, k int) int {
	freq := make(map[int]bool)
	for _, v := range nums {
		if v < k {
			return -1
		} else if v > k {
			freq[v] = true
		}
	}
	return len(freq)
}
