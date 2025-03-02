package merge_two_2d_arrays_by_summing_values

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	res := make([][]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] == nums2[j][0] {
			res = append(res, []int{nums1[i][0], nums2[j][1] + nums1[i][1]})
			i++
			j++
		} else if nums1[i][0] > nums2[j][0] {
			res = append(res, nums2[j])
			j++
		} else {
			res = append(res, nums1[i])
			i++
		}
	}
	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	}
	if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}
	return res
}
