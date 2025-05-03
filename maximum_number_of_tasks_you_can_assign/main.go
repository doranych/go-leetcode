package maximum_number_of_tasks_you_can_assign

import (
	"container/list"
	"sort"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	check := func(mid int) bool {
		p := pills
		ws := list.New()
		ptr := m - 1

		for i := mid - 1; i >= 0; i-- {
			for ptr >= m-mid && workers[ptr]+strength >= tasks[i] {
				ws.PushFront(workers[ptr])
				ptr--
			}
			if ws.Len() == 0 {
				return false
			}

			if ws.Back().Value.(int) >= tasks[i] {
				ws.Remove(ws.Back())
			} else {
				if p == 0 {
					return false
				}
				p--
				ws.Remove(ws.Front())
			}
		}
		return true
	}

	left, right, ans := 1, min(m, n), 0
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
