package maximum_score_after_splitting_a_string

func maxScore(s string) int {
	n := len(s)
	zero := []byte("0")[0]
	one := []byte("1")[0]
	sumLeft := make([]int, n)
	sumRight := make([]int, n)

	if s[0] == zero {
		sumLeft[0] = 1
	}
	if s[n-1] == one {
		sumRight[n-1] = 1
	}

	for i, j := 1, n-2; i < n-1 && j > 0; i, j = i+1, j-1 {
		if s[i] == zero {
			sumLeft[i] = sumLeft[i-1] + 1
		} else {
			sumLeft[i] = sumLeft[i-1]
		}

		if s[j] == one {
			sumRight[j] = sumRight[j+1] + 1
		} else {
			sumRight[j] = sumRight[j+1]
		}
	}

	sumRight[0] = sumRight[1]
	sumLeft[n-1] = sumLeft[n-2]

	maxSum := 0
	for i := 0; i < n; i++ {
		sum := sumLeft[i] + sumRight[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
