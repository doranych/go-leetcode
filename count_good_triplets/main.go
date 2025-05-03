package count_good_triplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	n := len(arr)
	for i := range n - 2 {
		for j := i + 1; j < n-1; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < n; k++ {
					if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
						res++
					}
				}
			}
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
