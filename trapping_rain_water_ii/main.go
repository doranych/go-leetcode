package trapping_rain_water_ii

import (
	"container/heap"
	"fmt"
)

/*
 https://leetcode.com/problems/trapping-rain-water-ii/description/
 Given an m x n integer matrix heightMap representing the height of each unit cell in a 2D elevation map,
 return the volume of water it can trap after raining.

 Example 1:
 Input: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
 Output: 4
 Explanation:
 After the rain, water is trapped between the blocks.
 We have two small pounds 1 and 3 units of water trapped.
 The total volume of water = 4 units.

 Example 2:
 Input: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
 Output: 10

 Constraints:
    m == heightMap.length
    n == heightMap[i].length
    1 <= m, n <= 200
    0 <= heightMap[i][j] <= 2 * 104
*/

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
	visited := make([][]bool, m)

	boundary := &MinHeap{}
	heap.Init(boundary)
	// create matrix of visited cells, push box to the boundary queue
	// and mark initial boundary as visited
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		if i == 0 || i == m-1 {
			for j := 0; j < n; j++ {
				visited[i][j] = true
				boundary.Push(Cell{heightMap[i][j], i, j})
			}
		} else {
			visited[i][0] = true
			visited[i][n-1] = true
			boundary.Push(Cell{heightMap[i][0], i, 0})
			boundary.Push(Cell{heightMap[i][n-1], i, n - 1})
		}
	}

	d_x := []int{0, 0, -1, 1}
	d_y := []int{-1, 1, 0, 0}

	totalWater := 0
	for boundary.Len() > 0 {
		cell := heap.Pop(boundary).(Cell)
		fmt.Println(cell)
		// looking around
		for i := range 4 {
			newX, newY := cell.x+d_x[i], cell.y+d_y[i]
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
