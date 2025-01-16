package bitwise_xor_of_all_pairings

func xorAllNums(nums1 []int, nums2 []int) int {
	r := 0
	if len(nums2)%2 == 1 {
		for _, v := range nums1 {
			r ^= v
		}
	}
	if len(nums1)%2 == 1 {
		for _, v := range nums2 {
			r ^= v
		}
	}

	return r
}
