package minimum_domino_rotations_for_equal_row

func minDominoRotations(A []int, B []int) int {
	n := len(A)
	check := func(x int) int {
		rotationsA, rotationsB := 0, 0
		for i := 0; i < n; i++ {
			if A[i] != x && B[i] != x {
				return -1
			}
			if A[i] != x {
				rotationsA++
			}
			if B[i] != x {
				rotationsB++
			}
		}
		return min(rotationsA, rotationsB)
	}

	rotations := check(A[0])
	if rotations != -1 || A[0] == B[0] {
		return rotations
	}
	return check(B[0])
}
