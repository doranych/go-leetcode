package length_of_longest_fibonacci_subsequence

func lenLongestFibSubseq(arr []int) int {
	numSet := make(map[int]struct{})
	for _, num := range arr {
		numSet[num] = struct{}{}
	}

	maxLen := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			prev := arr[j]
			curr := arr[i] + arr[j]
			currLen := 2

			for _, ok := numSet[curr]; ok; _, ok = numSet[curr] {
				prev, curr = curr, curr+prev
				currLen++
				if currLen > maxLen {
					maxLen = currLen
				}
			}
		}
	}

	return maxLen
}
