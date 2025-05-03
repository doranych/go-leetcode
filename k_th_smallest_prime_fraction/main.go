package k_th_smallest_prime_fraction

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	if len(arr) == 2 {
		return arr
	}
	m := make(map[float64][]int)
	tempRes := make([]float64, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			v := float64(arr[i]) / float64(arr[j])
			tempRes = append(tempRes, v)
			m[v] = []int{arr[i], arr[j]}
		}
	}
	sort.Float64s(tempRes)
	return m[tempRes[k-1]]
}
