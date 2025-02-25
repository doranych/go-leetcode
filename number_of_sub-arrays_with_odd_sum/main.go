package number_of_sub_arrays_with_odd_sum

var mod = int(1e9 + 7)

func numOfSubarrays(arr []int) int {
	sum, prefixSum, odd, even := 0, 0, 0, 1
	for _, v := range arr {
		prefixSum += v
		if prefixSum%2 == 0 {
			sum += odd
			even++
		} else {
			sum += even
			odd++
		}
		sum = sum % mod
	}
	return sum
}
