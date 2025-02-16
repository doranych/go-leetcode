package construct_the_lexicographically_largest_valid_sequence

var (
	seq  []int
	used []bool
)

func constructDistancedSequence(n int) []int {
	seq = make([]int, 2*n-1)
	used = make([]bool, n+1)

	findSeq(0, n)
	return seq
}

func findSeq(curIdx, n int) bool {
	if curIdx == len(seq) {
		return true
	}
	if seq[curIdx] != 0 {
		return findSeq(curIdx+1, n)
	}

	for i := n; i > 0; i-- {
		if used[i] {
			continue
		}
		used[i] = true
		seq[curIdx] = i

		if i == 1 {
			if findSeq(curIdx+1, n) {
				return true
			}
		} else if i+curIdx < len(seq) && seq[i+curIdx] == 0 {
			seq[i+curIdx] = i

			if findSeq(curIdx+1, n) {
				return true
			}
			seq[curIdx+i] = 0
		}
		seq[curIdx] = 0
		used[i] = false
	}
	return false
}
