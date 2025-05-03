package count_subarrays_of_length_three_with_a_condition

func countSubarrays(nums []int) int {
	cnt := 0
	for i := 2; i < len(nums); i++ {
		if (nums[i-2]+nums[i])*2 == nums[i-1] {
			cnt++
		}
	}
	return cnt
}
