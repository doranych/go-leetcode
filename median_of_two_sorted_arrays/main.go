package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	nums := join(nums1, nums2)
	if len(nums)%2 == 0 {
		a := len(nums) / 2
		b := a - 1
		return float64(nums[a]+nums[b]) / 2
	} else {
		return float64(nums[len(nums)/2])
	}
}

func join(nums1 []int, nums2 []int) []int {
	nums := make([]int, 0)
	a := 0
	b := 0
	for {
		if a < len(nums1) && b < len(nums2) {
			if nums1[a] < nums2[b] {
				nums = append(nums, nums1[a])
				a++
			} else {
				nums = append(nums, nums2[b])
				b++
			}
		} else if a < len(nums1) {
			nums = append(nums, nums1[a])
			a++
		} else if b < len(nums2) {
			nums = append(nums, nums2[b])
			b++
		}
		if a >= len(nums1) && b >= len(nums2) {
			return nums
		}
	}
}
