package kids_with_the_greatest_number_of_candies

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxAmount := slices.Max(candies)
	res := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= maxAmount {
			res[i] = true
			continue
		}
		res[i] = false
	}
	return res
}
