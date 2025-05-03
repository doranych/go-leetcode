package count_of_interesting_subarrays

func countInterestingSubarrays(nums []int, mod int, k int) int64 {
	n := len(nums)
	cnt := make(map[int]int)
	var res int64 = 0
	var p = 0
	cnt[0] = 1
	for i := 0; i < n; i++ {
		if nums[i]%mod == k {
			p++
		}
		res += int64(cnt[(p-k+mod)%mod])
		cnt[p%mod]++
	}
	return res
}
