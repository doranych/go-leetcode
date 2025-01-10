package zigzag_conversion

func convert(s string, numRows int) string {
	if numRows >= len(s) || numRows == 1 {
		return s
	}
	m := make([][]byte, numRows)

	j := 0
	r := 0
	dir := 0
	for i := 0; i < len(s); i++ {
		if m[r] == nil {
			m[r] = make([]byte, len(s)/2+1)
		}
		m[r][j] = s[i]
		if dir == 0 {
			if r < numRows-1 {
				r++
			} else {
				r--
				j++
				dir = 1
			}
		} else {
			if r > 0 {
				r--
				j++
			} else {
				dir = 0
				r++
			}
		}
	}
	res := make([]byte, len(s))

	i := 0
	for r := 0; r < numRows; r++ {
		for j := 0; j < len(s)/2+1; j++ {
			if m[r][j] != 0 {
				res[i] = m[r][j]
				i++
				if i >= len(s) {
					break
				}
			}
		}
	}
	return string(res)
}
