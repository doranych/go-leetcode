package maximum_value_of_an_ordered_triplet_ii

func maximumTripletValue(nums []int) int64 {
	res, i, d := 0, 0, 0
	for k := 0; k < len(nums); k++ {
		res = max(res, d*nums[k])
		d = max(d, i-nums[k])
		i = max(i, nums[k])
	}
	return int64(res)
}
