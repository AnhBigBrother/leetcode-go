package main

import "container/heap"

func minimumTime2577(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	time := [][]int{}
	for i := 0; i < m; i++ {
		row := []int{}
		for i := 0; i < n; i++ {
			row = append(row, 1e9)
		}
		time = append(time, row)
	}

	time[0][0] = 0
	h := PriorityQueue{}
	heap.Init(&h)
	heap.Push(&h, &PQItem{
		Value:    0,
		Priority: 0,
	})

	for len(h) > 0 {
		p := heap.Pop(&h).(*PQItem).Value
		x, y := p/n, p%n
		if x == m-1 && y == n-1 {
			return time[x][y]
		}
		for _, k := range directions {
			nx, ny := x+k[0], y+k[1]
			if nx == -1 || nx == m || ny == -1 || ny == n {
				continue
			}
			newTime := int(1e9)
			if time[x][y] >= grid[nx][ny] {
				newTime = time[x][y] + 1
			} else {
				diff := grid[nx][ny] - time[x][y]
				if diff%2 == 1 {
					newTime = grid[nx][ny]
				} else {
					newTime = grid[nx][ny] + 1
				}
			}
			if newTime < time[nx][ny] {
				time[nx][ny] = newTime
				heap.Push(&h, &PQItem{
					Value:    nx*n + ny,
					Priority: newTime,
				})
			}
		}
	}

	return -1
}
