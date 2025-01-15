package minimize_xor

import "math/bits"

func minimizeXor(num1 int, num2 int) int {
	res := num1
	resultBits := bits.OnesCount(uint(res))
	bitsToHave := bits.OnesCount(uint(num2))

	currentBit := 0
	for bitsToHave > resultBits {
		if res&(1<<currentBit) == 0 {
			res |= 1 << currentBit
			resultBits++
		}
		currentBit++
	}
	currentBit = 0
	for bitsToHave < resultBits {
		if res&(1<<currentBit) != 0 {
			res &= ^(1 << currentBit)
			resultBits--
		}
		currentBit++
	}
	return res
}
