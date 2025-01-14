package find_the_prefix_common_array_of_two_arrays

func findThePrefixCommonArray(A []int, B []int) []int {
	counterAB := make([]int, len(A))
	c := make([]int, len(A))

	sum := 0
	for i := 0; i < len(A); i++ {
		counterAB[A[i]-1]++
		counterAB[B[i]-1]++
		if A[i] == B[i] {
			sum++
		} else {
			if counterAB[A[i]-1] == 2 {
				sum++
			}
			if A[i] != B[i] && counterAB[B[i]-1] == 2 {
				sum++
			}
		}
		c[i] = sum
	}
	return c
}
