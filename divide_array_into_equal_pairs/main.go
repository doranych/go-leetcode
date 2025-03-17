package divide_array_into_equal_pairs

func divideArray(nums []int) bool {
	fr := make(map[int]uint, len(nums)/2)
	for i := range nums {
		fr[nums[i]]++
	}
	for _, v := range fr {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
