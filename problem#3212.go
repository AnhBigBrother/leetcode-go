package main

func numberOfSubmatrices(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	countX := [][]int{}
	countY := [][]int{}
	for i := 0; i < m; i++ {
		countX = append(countX, make([]int, n))
		countY = append(countY, make([]int, n))
	}
	for i := 0; i < m; i++ {
		x := 0
		y := 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 'X' {
				x++
			}
			if grid[i][j] == 'Y' {
				y++
			}
			countX[i][j] = x
			countY[i][j] = y
		}
	}

	ans := 0
	for j := 0; j < n; j++ {
		x := countX[0][j]
		y := countY[0][j]
		if x > 0 && x == y {
			ans++
		}
		for i := 1; i < m; i++ {
			x += countX[i][j]
			y += countY[i][j]
			if x > 0 && x == y {
				ans++
			}
		}
	}

	return ans
}
