package main

type coor struct {
	x int
	y int
}

func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := [][]bool{}
	island := [][]int{}
	square := map[int]int{}
	for i := 0; i < m; i++ {
		island = append(island, make([]int, n))
		visited = append(visited, make([]bool, n))
	}
	bfs := func(a, b, isl int) {
		visited[a][b] = true
		q := []coor{{a, b}}
		for i := 0; i < len(q); i++ {
			x, y := q[i].x, q[i].y
			island[x][y] = isl
			if x-1 >= 0 && grid[x-1][y] > 0 && !visited[x-1][y] {
				visited[x-1][y] = true
				q = append(q, coor{x - 1, y})
			}
			if x+1 < m && grid[x+1][y] > 0 && !visited[x+1][y] {
				visited[x+1][y] = true
				q = append(q, coor{x + 1, y})
			}
			if y-1 >= 0 && grid[x][y-1] > 0 && !visited[x][y-1] {
				visited[x][y-1] = true
				q = append(q, coor{x, y - 1})
			}
			if y+1 < n && grid[x][y+1] > 0 && !visited[x][y+1] {
				visited[x][y+1] = true
				q = append(q, coor{x, y + 1})
			}
		}
		square[isl] = len(q)
	}

	isl := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] > 0 {
				isl++
				bfs(i, j, isl)
			}
		}
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, square[island[i][j]])
			if grid[i][j] == 0 {
				adjacent_ilsand := map[int]bool{}
				if i-1 >= 0 {
					adjacent_ilsand[island[i-1][j]] = true
				}
				if i+1 < m {
					adjacent_ilsand[island[i+1][j]] = true
				}
				if j-1 >= 0 {
					adjacent_ilsand[island[i][j-1]] = true
				}
				if j+1 < n {
					adjacent_ilsand[island[i][j+1]] = true
				}
				s := 1
				for i := range adjacent_ilsand {
					s += square[i]
				}
				ans = max(ans, s)
			}
		}
	}

	return ans
}
