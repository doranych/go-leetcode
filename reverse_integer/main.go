package reverse_integer

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	s := strconv.Itoa(x)
	rev := make([]byte, len(s))
	cur := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			rev[cur] = s[i]
			cur++
			continue
		}

		rev = append([]byte{'-'}, rev[:cur]...)
	}
	fmt.Println(string(rev))
	v, _ := strconv.Atoi(string(rev))
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0
	}
	return v
}
