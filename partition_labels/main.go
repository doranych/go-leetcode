package partition_labels

func partitionLabels(s string) []int {
	p := [26]int{}
	for i := range s {
		p[s[i]-'a'] = i
	}
	i := 0
	res := make([]int, 0)
	segment := 0
	start := 0
	for i < len(s) {
		if p[s[i]-'a'] > segment {
			segment = p[s[i]-'a']
		}
		if segment == i {
			res = append(res, segment-start+1)
			start = i + 1
		}
		i++
	}
	return res
}
