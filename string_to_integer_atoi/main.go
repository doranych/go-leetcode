package string_to_integer_atoi

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	positive := true
	leading := true
	sawDigit := false
	number := make([]uint8, 0)

loop:
	for i := 0; i < len(s); i++ {
		switch true {
		case s[i] == ' ':
			if !sawDigit {
				continue
			}
			break loop
		case s[i] == '.' || s[i] > '9':
			break loop
		case s[i] == '+':
			if sawDigit {
				break loop
			}
			sawDigit = true
		case s[i] == '-':
			if sawDigit {
				break loop
			}
			sawDigit = true
			positive = false
		case s[i] == '0':
			sawDigit = true
			if leading {
				continue
			}
			number = append(number, s[i])
		case s[i] > '0' && s[i] <= '9':
			if !sawDigit {
				sawDigit = true
			}
			leading = false
			number = append(number, s[i])
		}
	}
	res := 0
	pos := 0
	for i := len(number) - 1; i >= 0; i-- {
		num := int(number[i] - '0')
		res += num * int(math.Pow10(pos))
		fmt.Println(res)
		if positive && (res > math.MaxInt32 || i > 10) {
			return math.MaxInt32
		}
		if !positive && (-res < math.MinInt32 || i > 10) {
			return math.MinInt32
		}
		pos++
	}
	if !positive {
		res = -res
	}
	return res
}
