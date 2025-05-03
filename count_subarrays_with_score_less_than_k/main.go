package count_subarrays_with_score_less_than_k

func countSubarrays(nums []int, k int64) int64 {
	l := 0
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum*(i-l+1) >= int(k) {
			sum -= nums[l]
			l++
		}
		res += i - l + 1
	}
	return int64(res)
}
