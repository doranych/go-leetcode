package count_the_number_of_powerful_integers

import (
	"math"
	"strconv"
)

func numberOfPowerfulInt(st int64, fin int64, l int, s string) int64 {
	return calculate(strconv.FormatInt(fin, 10), s, l) - calculate(strconv.FormatInt(st-1, 10), s, l)
}

func calculate(x, s string, l int) int64 {
	if len(x) < len(s) {
		return 0
	}
	if len(x) == len(s) {
		if x >= s {
			return 1
		}
		return 0
	}

	preLen := len(x) - len(s)
	suf := x[preLen:]
	var c int64

	for i := 0; i < preLen; i++ {
		digit := int(x[i] - '0')
		if l < digit {
			c += int64(math.Pow(float64(l+1), float64(preLen-i)))
			return c
		}
		c += int64(digit) * int64(math.Pow(float64(l+1), float64(preLen-1-i)))
	}
	if suf >= s {
		c++
	}
	return c
}
