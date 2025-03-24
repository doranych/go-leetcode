package count_days_without_meetings

import (
	"maps"
	"slices"
)

func countDays(days int, meetings [][]int) int {
	sweep := make(map[int]int, 0)
	pDay := days
	for _, r := range meetings {
		pDay = min(pDay, r[0])
		sweep[r[0]]++
		sweep[r[1]+1]--
	}
	count := pDay - 1
	sum := 0
	for _, i := range slices.Sorted(maps.Keys(sweep)) {
		if sum == 0 {
			count += i - pDay
		}
		sum += sweep[i]
		pDay = i
	}
	count += days - pDay + 1
	return count
}
