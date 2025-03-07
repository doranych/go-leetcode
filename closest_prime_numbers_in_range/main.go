package closest_prime_numbers_in_range

import "math"

func sieve(limit int) []bool {
	list := make([]bool, limit+1)
	for i := range list {
		list[i] = true
	}
	list[0], list[1] = false, false

	for n := 2; n <= int(math.Sqrt(float64(limit))); n++ {
		if list[n] {
			for m := n * n; m <= limit; m += n {
				list[m] = false
			}
		}
	}
	return list
}

func closestPrimes(left, right int) []int {
	list := sieve(right)

	primes := []int{}
	for num := left; num <= right; num++ {
		if list[num] {
			primes = append(primes, num)
		}
	}

	pair := []int{-1, -1}

	if len(primes) < 2 {
		return pair
	}
	minDifference := math.MaxInt32

	for i := 1; i < len(primes); i++ {
		diff := primes[i] - primes[i-1]
		if diff < minDifference {
			minDifference = diff
			pair = []int{primes[i-1], primes[i]}
		}
	}

	return pair
}
