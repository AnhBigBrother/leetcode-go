package main

import "container/heap"

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	minObs := [][]int{}
	for i := 0; i < m; i++ {
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, 100001)
		}
		minObs = append(minObs, row)
	}
	minObs[0][0] = 0

	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &PQItem{
		Value:    0,
		Priority: 0,
	})

	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*PQItem).Value
		x, y := p/n, p%n
		for _, k := range dir {
			nx, ny := x+k[0], y+k[1]
			if nx == -1 || nx == m || ny == -1 || ny == n {
				continue
			}
			if minObs[nx][ny] > minObs[x][y]+grid[x][y] {
				minObs[nx][ny] = minObs[x][y] + grid[x][y]
				heap.Push(&pq, &PQItem{
					Value:    nx*n + ny,
					Priority: minObs[nx][ny],
				})
			}
		}
	}

	return minObs[m-1][n-1]
}
