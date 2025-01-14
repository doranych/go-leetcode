package find_all_numbers_disappeared_in_an_array

import "math"

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := int(math.Abs(float64(nums[i])))
		nums[n-1] = -int(math.Abs(float64(nums[n-1])))
	}

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
