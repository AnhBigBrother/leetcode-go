package main

type coordinates1368 struct {
	x int
	y int
}

func minCost1368(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cost := [][]int{}
	for i := 0; i < m; i++ {
		row := []int{}
		for i := 0; i < n; i++ {
			row = append(row, 1000000)
		}
		cost = append(cost, row)
	}
	cost[0][0] = 0
	arr := []coordinates1368{{0, 0}}
	for len(arr) > 0 {
		x, y := arr[0].x, arr[0].y
		arr = arr[1:]
		if x+1 < m {
			k := 0
			if grid[x][y] != 3 {
				k = 1
			}
			if cost[x+1][y] > cost[x][y]+k {
				cost[x+1][y] = cost[x][y] + k
				arr = append(arr, coordinates1368{x + 1, y})
			}
		}
		if x-1 >= 0 {
			k := 0
			if grid[x][y] != 4 {
				k = 1
			}
			if cost[x-1][y] > cost[x][y]+k {
				cost[x-1][y] = cost[x][y] + k
				arr = append(arr, coordinates1368{x - 1, y})
			}
		}
		if y+1 < n {
			k := 0
			if grid[x][y] != 1 {
				k = 1
			}
			if cost[x][y+1] > cost[x][y]+k {
				cost[x][y+1] = cost[x][y] + k
				arr = append(arr, coordinates1368{x, y + 1})
			}
		}
		if y-1 >= 0 {
			k := 0
			if grid[x][y] != 2 {
				k = 1
			}
			if cost[x][y-1] > cost[x][y]+k {
				cost[x][y-1] = cost[x][y] + k
				arr = append(arr, coordinates1368{x, y - 1})
			}
		}
	}
	return cost[m-1][n-1]
}
