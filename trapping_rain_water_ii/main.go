package trapping_rain_water_ii

import (
	"container/heap"
)

type Cell struct {
	curHeight, x, y int
}

type MinHeap []Cell

// implement heap.Interface
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].curHeight < h[j].curHeight }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(c any) {
	*h = append(*h, c.(Cell))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	// create matrix of visited cells, push box to the boundary queue
	// and mark initial boundary as visited
	visited := make([][]bool, m)
	boundary := &MinHeap{}
	heap.Init(boundary)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				visited[i][j] = true
				heap.Push(boundary, Cell{heightMap[i][j], i, j})
			}
		}
	}

	dX := []int{0, 0, -1, 1}
	dY := []int{-1, 1, 0, 0}

	totalWater := 0
	for boundary.Len() > 0 {
		cell := heap.Pop(boundary).(Cell)
		// looking around
		for i := range 4 {
			newX, newY := cell.x+dX[i], cell.y+dY[i]
			if newX < 0 || newY < 0 || newX >= m || newY >= n {
				continue
			}
			if visited[newX][newY] {
				continue
			}

			if heightMap[newX][newY] < cell.curHeight {
				totalWater += cell.curHeight - heightMap[newX][newY]
			}

			visited[newX][newY] = true
			heap.Push(boundary, Cell{max(cell.curHeight, heightMap[newX][newY]), newX, newY})
		}
	}
	return totalWater
}
