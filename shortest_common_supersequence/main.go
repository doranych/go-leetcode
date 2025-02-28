package shortest_common_supersequence

func shortestCommonSupersequence(str1, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)

	prevRow := make([]string, n2+1)
	for col := 0; col <= n2; col++ {
		prevRow[col] = str2[:col]
	}

	for row := 1; row <= n1; row++ {
		currRow := make([]string, n2+1)
		currRow[0] = str1[:row]

		for col := 1; col <= n2; col++ {
			if str1[row-1] == str2[col-1] {
				currRow[col] = prevRow[col-1] + string(str1[row-1])
			} else {
				pickS1 := prevRow[col]
				pickS2 := currRow[col-1]

				if len(pickS1)+1 < len(pickS2)+1 {
					currRow[col] = pickS1 + string(str1[row-1])
				} else {
					currRow[col] = pickS2 + string(str2[col-1])
				}
			}
		}
		prevRow = currRow
	}

	return prevRow[n2]
}
