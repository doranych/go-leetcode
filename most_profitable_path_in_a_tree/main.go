package most_profitable_path_in_a_tree

import "math"

var tree map[int][]int
var distances []int

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	tree = make(map[int][]int, len(amount))
	distances = make([]int, len(amount))
	for _, p := range edges {
		tree[p[0]] = append(tree[p[0]], p[1])
		tree[p[1]] = append(tree[p[1]], p[0])
	}
	return dfs(0, 0, 0, bob, amount)
}

func dfs(src, par, time, bob int, amount []int) int {
	maxI := 0
	maxCh := math.MinInt

	if src == bob {
		distances[src] = 0
	} else {
		distances[src] = len(amount)
	}

	for _, node := range tree[src] {
		if node == par {
			continue
		}
		maxCh = max(maxCh, dfs(node, src, time+1, bob, amount))
		distances[src] = min(distances[src], distances[node]+1)
	}

	if distances[src] > time {
		maxI += amount[src]
	} else if distances[src] == time {
		maxI += amount[src] / 2
	}

	if maxCh == math.MinInt {
		return maxI
	}
	return maxI + maxCh
}
