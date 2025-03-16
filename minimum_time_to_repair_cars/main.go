package minimum_time_to_repair_cars

import (
	"container/heap"
	"slices"
)

func repairCars(ranks []int, cars int) int64 {
	counter := make([]int, slices.Max(ranks)+1)
	for i := range ranks {
		counter[ranks[i]]++
	}
	h := &minHeap{}
	heap.Init(h)
	for i := range counter {
		if counter[i] > 0 {
			heap.Push(h, repair{i, i, 1, counter[i]})
		}
	}
	t := 0
	for cars > 0 {
		var r, n, c int
		cur := heap.Pop(h).(repair)
		t, r, n, c = cur.time, cur.rank, cur.n, cur.count
		cars -= c
		n++
		heap.Push(h, repair{r * n * n, r, n, c})
	}
	return int64(t)
}

type repair struct {
	time  int
	rank  int
	n     int
	count int
}

type minHeap []repair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(p any) {
	*h = append(*h, p.(repair))
}
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
