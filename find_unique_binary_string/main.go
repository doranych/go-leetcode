package find_unique_binary_string

import "strconv"

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	maxN := 0
	for i := 0; i < n; i++ {
		maxN |= 1 << i
	}
	ns := make([]int, maxN+1)
	for _, nm := range nums {
		num, _ := strconv.ParseInt(nm, 2, 64)
		ns[num] = 1
	}
	for i := 0; i <= maxN; i++ {
		if ns[i] == 0 {
			s := strconv.FormatInt(int64(i), 2)
			if len(s) < n {
				for i := len(s); i < n; i++ {
					s = "0" + s
				}
			}
			return s
		}
	}
	return ""
}
