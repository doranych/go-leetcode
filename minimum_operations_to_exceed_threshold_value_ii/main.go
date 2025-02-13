package minimum_operations_to_exceed_threshold_value_ii

import "container/heap"

func minOperations(nums []int, k int) int {
	q := minQ(nums)
	heap.Init(&q)
	counter := 0
	for true {
		v1 := heap.Pop(&q).(int)
		if v1 >= k {
			break
		}
		v2 := heap.Pop(&q).(int)
		heap.Push(&q, v1*2+v2)
		counter++
	}
	return counter
}

type minQ []int

func (m minQ) Len() int           { return len(m) }
func (m minQ) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m minQ) Less(i, j int) bool { return m[i] < m[j] }
func (m *minQ) Push(n any) {
	*m = append(*m, n.(int))
}
func (m *minQ) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
