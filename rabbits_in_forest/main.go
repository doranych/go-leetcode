package rabbits_in_forest

import "math"

func numRabbits(answers []int) int {
	count := make(map[int]int)
	for _, answer := range answers {
		count[answer]++
	}

	totalRabbits := 0
	for k, v := range count {
		groupSize := k + 1
		groups := int(math.Ceil(float64(v) / float64(groupSize)))
		totalRabbits += groups * groupSize
	}

	return totalRabbits
}
