package count_number_of_bad_pairs

func countBadPairs(nums []int) int64 {
	res := 0
	diff := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		d := nums[i] - i
		res += i - diff[d]
		diff[d]++
	}
	return int64(res)
}
