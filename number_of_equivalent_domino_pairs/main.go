package number_of_equivalent_domino_pairs

func numEquivDominoPairs(dominoes [][]int) int {
	ans := 0
	cnt := [100]int{}
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return ans
}
