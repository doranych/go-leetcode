package container_with_most_water

func maxArea(height []int) int {
	sq := 0
	i, j := 0, len(height)-1
	for i < j {
		h := min(height[i], height[j])
		csq := (j - i) * h
		if csq > sq {
			sq = csq
		}
		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}
	return sq
}
