package count_nice_pairs_in_an_array

func countNicePairs(nums []int) int {
	res := 0
	g := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		d := nums[i] - rev(nums[i])
		res += g[d]
		g[d]++
	}
	return res % (1e9 + 7)
}

func rev(num int) int {
	if num < 9 {
		return num
	}
	var res int
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}
