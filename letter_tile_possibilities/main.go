package letter_tile_possibilities

func numTilePossibilities(tiles string) int {
	freq := [26]int{}
	for i := 0; i < len(tiles); i++ {
		freq[tiles[i]-'A'] += 1
	}
	return counter(freq)
}

func counter(freq [26]int) int {
	res := 0
	for i := range freq {
		if freq[i] == 0 {
			continue
		}

		res++
		freq[i]--
		res += counter(freq)
		freq[i]++
	}
	return res
}
