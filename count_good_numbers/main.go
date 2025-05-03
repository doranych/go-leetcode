package count_good_numbers

const mod = 1e9 + 7

func countGoodNumbers(n int64) int {
	mult := func(x, y int64) int64 {
		res := int64(1)
		mul := x
		for y > 0 {
			if y%2 == 1 {
				res = (res * mul) % mod
			}
			mul = (mul * mul) % mod
			y /= 2
		}
		return res
	}
	return int((mult(5, (n+1)/2) * mult(4, n/2)) % mod)
}
