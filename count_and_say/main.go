package count_and_say

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = rle(res)
	}
	return res
}

func rle(s string) string {
	sb := strings.Builder{}
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			sb.WriteString(fmt.Sprintf("%d%c", count, s[i-1]))
			count = 1
		}
	}
	sb.WriteString(fmt.Sprintf("%d%c", count, s[len(s)-1]))
	return sb.String()
}
