package number_of_ways_to_arrive_at_destination

import (
	"container/heap"
	"math"
)

const mod = 1e9 + 7

func countPaths(n int, roads [][]int) int {
	g := make([][][2]int, n)
	dist := make([]int, n)
	paths := make([]int, n)
	for i := range n {
		dist[i] = math.MaxInt64
	}
	for _, r := range roads {
		g[r[0]] = append(g[r[0]], [2]int{r[1], r[2]})
		g[r[1]] = append(g[r[1]], [2]int{r[0], r[2]})
	}

	paths[0] = 1
	dist[0] = 0
	h := &minHeap{node{0, 0}}
	heap.Init(h)

	for h.Len() > 0 {
		curN := heap.Pop(h).(node)
		cT, cN := curN[0], curN[1]
		if cT > dist[cN] {
			continue
		}
		for _, neigh := range g[cN] {
			nN, nT := neigh[0], neigh[1]
			if cT+nT < dist[nN] {
				dist[nN] = cT + nT
				heap.Push(h, node{dist[nN], nN})
				paths[nN] = paths[cN]
			} else if cT+nT == dist[nN] {
				paths[nN] = (paths[cN] + paths[nN]) % mod
			}
		}

	}

	return paths[n-1]
}

type (
	node    [2]int
	minHeap []node
)

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(r any) {
	*h = append(*h, r.(node))
}
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
