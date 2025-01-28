package main

type coor2658 struct {
	x int
	y int
}

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				sum := 0
				q := []coor2658{{i, j}}
				for k := 0; k < len(q); k++ {
					x, y := q[k].x, q[k].y
					sum += grid[x][y]
					grid[x][y] = 0
					if x-1 >= 0 && grid[x-1][y] > 0 {
						q = append(q, coor2658{x - 1, y})
					}
					if x+1 < m && grid[x+1][y] > 0 {
						q = append(q, coor2658{x + 1, y})
					}
					if y-1 >= 0 && grid[x][y-1] > 0 {
						q = append(q, coor2658{x, y - 1})
					}
					if y+1 < n && grid[x][y+1] > 0 {
						q = append(q, coor2658{x, y + 1})
					}
				}
				ans = max(ans, sum)
			}
		}
	}
	return ans
}
