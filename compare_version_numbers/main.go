package compare_version_numbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	reversed := false
	if len(v1) < len(v2) {
		v1, v2 = v2, v1
		reversed = true
	}

	i, j := 0, 0
	compareRes := 0
	for {
		if i >= len(v1) && j >= len(v2) || compareRes != 0 {
			break
		}
		v2Pos := 0
		if j < len(v2) {
			v2Pos, _ = strconv.Atoi(v2[j])
			j++
		}
		v1Pos := 0
		if i < len(v1) {
			v1Pos, _ = strconv.Atoi(v1[i])
			i++
		}

		if v1Pos > v2Pos {
			compareRes = 1
		}
		if v2Pos > v1Pos {
			compareRes = -1
		}
	}

	if reversed {
		return -(compareRes)
	}
	return compareRes
}
