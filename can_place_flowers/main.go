package can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	cnt := 0
	res := 0
	for i, v := range flowerbed {
		if i == 0 {
			cnt++
		}
		if v == 0 {
			cnt++
		} else if v == 1 {
			res += (cnt - 1) / 2
			cnt = 0
		}
	}
	if cnt >= 2 {
		res += cnt / 2
	}
	return res >= n
}
