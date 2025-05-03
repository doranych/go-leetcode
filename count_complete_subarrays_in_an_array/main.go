package count_complete_subarrays_in_an_array

func countCompleteSubarrays(nums []int) int {
	res := 0
	cnt := make(map[int]int)
	dNums := make(map[int]bool)
	for _, num := range nums {
		dNums[num] = true
	}
	r := 0
	for l := 0; l < len(nums); l++ {
		if l > 0 {
			cnt[nums[l-1]]--
			if cnt[nums[l-1]] == 0 {
				delete(cnt, nums[l-1])
			}
		}
		for r < len(nums) && len(cnt) < len(dNums) {
			add := nums[r]
			cnt[add]++
			r++
		}
		if len(cnt) == len(dNums) {
			res += (len(nums) - r + 1)
		}
	}
	return res
}
