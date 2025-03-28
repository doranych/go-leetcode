package maximum_number_of_points_from_grid_queries

import "sort"

var dir = []int{0, 1, 0, -1, 0}

func maxPoints(grid [][]int, queries []int) []int {
	rc, cc := len(grid), len(grid[0])
	tc := rc * cc

	q := make([]query, len(queries))
	for i, val := range queries {
		q[i] = query{i, val}
	}
	sort.Slice(q, func(i, j int) bool { return q[i].value < q[j].value })

	c := make([]cell, tc)
	idx := 0
	for i := 0; i < rc; i++ {
		for j := 0; j < cc; j++ {
			c[idx] = cell{i, j, grid[i][j]}
			idx++
		}
	}
	sort.Slice(c, func(i, j int) bool { return c[i].value < c[j].value })

	u := &uf{
		parent: make([]int, tc),
		size:   make([]int, tc),
	}
	for i := range u.parent {
		u.parent[i] = -1
		u.size[i] = 1
	}
	res := make([]int, len(queries))

	cI := 0
	for _, qry := range q {
		for cI < tc && c[cI].value < qry.value {
			row := c[cI].row
			col := c[cI].col
			cId := row*cc + col

			for i := 0; i < 4; i++ {
				nr := row + dir[i]
				nc := col + dir[i+1]
				if nr >= 0 && nr < rc && nc >= 0 && nc < cc && grid[nr][nc] < qry.value {
					u.union(cId, nr*cc+nc)
				}
			}
			cI++
		}
		if qry.value > grid[0][0] {
			res[qry.index] = u.getSize(0)
		} else {
			res[qry.index] = 0
		}
	}

	return res
}

type cell struct {
	row, col, value int
}

type query struct {
	index, value int
}

type uf struct {
	parent, size []int
}

func (u *uf) find(node int) int {
	if u.parent[node] < 0 {
		return node
	}
	u.parent[node] = u.find(u.parent[node])
	return u.parent[node]
}

func (u *uf) union(nodeA, nodeB int) bool {
	ra := u.find(nodeA)
	rb := u.find(nodeB)
	if ra == rb {
		return false
	}

	if u.size[ra] > u.size[rb] {
		u.parent[rb] = ra
		u.size[ra] += u.size[rb]
	} else {
		u.parent[ra] = rb
		u.size[rb] += u.size[ra]
	}
	return true
}

func (u *uf) getSize(node int) int {
	return u.size[u.find(node)]
}
