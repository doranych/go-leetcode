package count_largest_group

func countLargestGroup(n int) int {
	m := make(map[int]int)
	maxG := 0
	for i := 1; i <= n; i++ {
		k := i
		c := 0
		for k > 0 {
			c += k % 10
			k /= 10
		}
		m[c]++
		maxG = max(maxG, m[c])
	}
	res := 0
	for _, v := range m {
		if v == maxG {
			res++
		}
	}
	return res
}
