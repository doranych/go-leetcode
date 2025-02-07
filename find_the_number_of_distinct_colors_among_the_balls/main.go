package find_the_number_of_distinct_colors_among_the_balls

func queryResults(_ int, queries [][]int) []int {
	bm := make(map[int]int)
	cm := make(map[int]int)
	res := make([]int, len(queries))
	for i, q := range queries {
		b, c := q[0], q[1]
		if tc, ok := bm[b]; ok {
			cm[tc]--
			if cm[tc] == 0 {
				delete(cm, tc)
			}
		}
		cm[c]++
		bm[b] = c
		res[i] = len(cm)
	}
	return res
}
