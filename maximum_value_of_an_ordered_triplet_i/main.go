package maximum_value_of_an_ordered_triplet_i

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	res, i, d := 0, 0, 0
	for k := 0; k < n; k++ {
		res = max(res, d*nums[k])
		d = max(d, i-nums[k])
		i = max(i, nums[k])
	}
	return int64(res)
}
