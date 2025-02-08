package design_a_number_container_system

import "sort"

type NumberContainers struct {
	values   map[int][]int
	indices  map[int]int
	isSorted map[int]bool
}

func Constructor() NumberContainers {
	return NumberContainers{
		values:   make(map[int][]int),
		indices:  make(map[int]int),
		isSorted: make(map[int]bool),
	}
}

func (nc *NumberContainers) Change(index int, number int) {
	if v, ok := nc.indices[index]; ok {
		if v != number {
			for i, n := range nc.values[v] {
				if n == index {
					nc.values[v] = append(nc.values[v][:i], nc.values[v][i+1:]...)
				}
			}
			delete(nc.indices, index)
			nc.isSorted[v] = false
			nc.indices[index] = number
			nc.values[number] = append(nc.values[number], index)
			nc.isSorted[number] = false
		}
	} else {
		nc.indices[index] = number
		nc.values[number] = append(nc.values[number], index)
		nc.isSorted[number] = false
	}
}

func (nc *NumberContainers) Find(number int) int {
	if len(nc.values[number]) == 0 {
		return -1
	}

	if !nc.isSorted[number] {
		sort.Ints(nc.values[number])
		nc.isSorted[number] = true
	}

	return nc.values[number][0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
