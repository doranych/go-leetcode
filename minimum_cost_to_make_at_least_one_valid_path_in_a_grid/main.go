package minimum_cost_to_make_at_least_one_valid_path_in_a_grid

import (
	"container/heap"
	"math"
)

type Cell struct {
	x, y, cost int
}
type MinHeap []Cell

// heap interfase
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Cell))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	costs := make([][]int, m)
	for i, _ := range costs {
		costs[i] = make([]int, n)
		for j, _ := range costs[i] {
			costs[i][j] = math.MaxInt32
		}
	}
	costs[0][0] = 0

	directions := []struct{ x, y int }{
		{0, 1},  // right
		{0, -1}, // left
		{1, 0},  // down
		{-1, 0}, // up
	}

	h := &MinHeap{}
	heap.Push(h, Cell{0, 0, 0})

	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		if cell.cost > costs[cell.x][cell.y] {
			continue
		}

		for i, _ := range directions {
			newX, newY := cell.x+directions[i].x, cell.y+directions[i].y
			if newX < 0 || newY < 0 || newX >= m || newY >= n {
				continue
			}

			newCost := cell.cost
			if grid[cell.x][cell.y] != i+1 {
				newCost++
			}

			if newCost < costs[newX][newY] {
				costs[newX][newY] = newCost
				heap.Push(h, Cell{newX, newY, newCost})
			}
		}
	}
	return costs[m-1][n-1]
}
