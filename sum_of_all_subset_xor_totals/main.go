package sum_of_all_subset_xor_totals

func subsetXORSum(nums []int) int {
	res := 0
	for i := range nums {
		res |= nums[i]
	}
	return res << (len(nums) - 1)
}
