package apply_operations_to_maximize_score

import (
	"container/heap"
	"math"
)

const mod = 1e9 + 7

func maximumScore(nums []int, k int) int {
	n := len(nums)
	scores := make([]int, n)
	next := make([]int, n)
	prev := make([]int, n)

	for i := range n {
		num := nums[i]
		for f := 2; f <= int(math.Sqrt(float64(num))); f++ {
			if num%f == 0 {
				scores[i]++
				for num%f == 0 {
					num /= f
				}
			}
		}
		if num >= 2 {
			scores[i]++
		}
		next[i] = n
		prev[i] = -1
	}

	stack := make([]int, 0)
	for i := range n {
		for len(stack) > 0 && scores[stack[len(stack)-1]] < scores[i] {
			tI := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next[tI] = i
		}

		if len(stack) > 0 {
			prev[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	numSubs := make([]int, n)
	for i := range n {
		numSubs[i] = (next[i] - i) * (i - prev[i])
	}

	h := &maxHeap{}
	heap.Init(h)
	for i := range n {
		heap.Push(h, [2]int{nums[i], i})
	}

	res := 1
	pow := func(b, e int) int {
		r := 1
		for e > 0 {
			if e%2 == 1 {
				r = (r * b) % mod
			}
			b = (b * b) % mod
			e /= 2
		}
		return r
	}

	for k > 0 {
		el := heap.Pop(h).([2]int)
		num, i := el[0], el[1]
		operations := min(k, numSubs[i])
		res = (res * pow(num, operations)) % mod
		k -= operations
	}

	return res
}

type maxHeap [][2]int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h *maxHeap) Push(el any) {
	*h = append(*h, el.([2]int))
}
func (h *maxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}
