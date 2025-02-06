package tuple_with_same_product

func tupleSameProduct(nums []int) int {
	prods := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			prods[nums[i]*nums[j]]++
		}
	}
	res := 0
	for _, v := range prods {
		if v > 1 {
			res += 4 * v * (v - 1)
		}
	}
	return res
}
