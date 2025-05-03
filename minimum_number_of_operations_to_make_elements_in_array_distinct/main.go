package minimum_number_of_operations_to_make_elements_in_array_distinct

func minimumOperations(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	c := 0
	for len(m) != len(nums) {
		if len(nums) == 0 {
			break
		}
		for _, v := range nums[:min(len(nums), 3)] {
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		}
		nums = nums[min(len(nums), 3):]
		c++
	}
	return c
}
