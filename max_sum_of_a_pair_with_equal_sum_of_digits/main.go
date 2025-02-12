package max_sum_of_a_pair_with_equal_sum_of_digits

func maximumSum(nums []int) int {
	res := -1
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		s := 0
		for n > 0 {
			s += n % 10
			n /= 10
		}
		localMax := -1
		if m[s] != 0 {
			localMax = nums[i] + m[s]
		}
		m[s] = max(m[s], nums[i])
		res = max(localMax, res)
	}
	return res
}
