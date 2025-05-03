package relative_ranks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	m := make(map[int]int, n)
	res := make([]string, n, n)
	for i, v := range score {
		m[v] = i
	}
	sort.Ints(score)
	curPos := 1
	for i := n - 1; i >= 0; i-- {
		v := score[i]
		res[m[v]] = posName(curPos)
		curPos++
	}
	return res
}

func posName(curPos int) string {
	switch curPos {
	case 1:
		return "Gold Medal"
	case 2:
		return "Silver Medal"
	case 3:
		return "Bronze Medal"
	default:
		return strconv.Itoa(curPos)
	}
}
