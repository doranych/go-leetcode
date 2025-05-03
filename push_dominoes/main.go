package push_dominoes

import "strings"

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	forces := make([]int, n)
	f := 0
	for i := range dominoes {
		if dominoes[i] == 'R' {
			f = n
		} else if dominoes[i] == 'L' {
			f = 0
		} else {
			f = max(f-1, 0)
		}
		forces[i] += f
	}
	f = 0
	for i := n - 1; i >= 0; i-- {
		if dominoes[i] == 'L' {
			f = n
		} else if dominoes[i] == 'R' {
			f = 0
		} else {
			f = max(f-1, 0)
		}
		forces[i] -= f
	}
	sb := strings.Builder{}
	for _, v := range forces {
		s := "."
		switch {
		case v < 0:
			s = "L"
		case v > 0:
			s = "R"
		}
		sb.WriteString(s)
	}
	return sb.String()
}
