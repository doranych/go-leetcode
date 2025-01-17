package neighboring_bitwise_xor

func doesValidArrayExist(derived []int) bool {
	s := 0
	for _, v := range derived {
		s ^= v
	}
	return s == 0
}
