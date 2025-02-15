package find_the_punishment_number_of_an_integer

func punishmentNumber(n int) int {
	sum := 0
	for i := range n {
		sq := (i + 1) * (i + 1)
		if canPartition(sq, i+1) {
			sum += sq
		}
	}
	return sum
}

func canPartition(n int, t int) bool {
	if t < 0 || n < t {
		return false
	}
	if n == t {
		return true
	}
	return canPartition(n/10, t-n%10) ||
		canPartition(n/100, t-n%100) ||
		canPartition(n/1000, t-n%1000)
}
