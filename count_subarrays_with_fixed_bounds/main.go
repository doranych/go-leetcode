package count_subarrays_with_fixed_bounds

func countSubarrays(nums []int, minK int, maxK int) int64 {
	answer := 0
	minPosition, maxPosition, leftBound := -1, -1, -1

	for i, number := range nums {
		if number < minK || number > maxK {
			leftBound = i
		}

		if number == minK {
			minPosition = i
		}
		if number == maxK {
			maxPosition = i
		}

		if minPosition != -1 && maxPosition != -1 {
			answer += max(0, min(minPosition, maxPosition)-leftBound)
		}
	}

	return int64(answer)
}
