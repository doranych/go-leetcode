package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	suf := make([]int, len(nums))
	res := make([]int, len(nums))
	n := len(nums)
	i, j := 0, n-1
	pre[0] = nums[i]
	suf[j] = nums[j]
	i++
	j--
	for {
		if i >= n && j <= 0 {
			break
		}
		pre[i] = nums[i] * pre[i-1]
		suf[j] = nums[j] * suf[j+1]
		i++
		j--
	}
	for i = 0; i < n; i++ {
		prefix := 1
		if i > 0 {
			prefix = pre[i-1]
		}
		suffix := 1
		if i < n-1 {
			suffix = suf[i+1]
		}
		res[i] = prefix * suffix
	}
	return res
}
