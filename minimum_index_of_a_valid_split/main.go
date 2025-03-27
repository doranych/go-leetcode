package minimum_index_of_a_valid_split

func minimumIndex(nums []int) int {
	m1, m2 := make(map[int]int), make(map[int]int)
	n := len(nums)
	for _, i := range nums {
		m2[i]++
	}

	for i, v := range nums {
		m1[v]++
		m2[v]--
		if m1[v]*2 > i+1 && m2[v]*2 > n-i-1 {
			return i
		}
	}
	return -1
}
