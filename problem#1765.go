package main

type coor1765 struct {
	x int
	y int
}

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	height := [][]int{}
	for i := 0; i < m; i++ {
		r := []int{}
		for j := 0; j < n; j++ {
			r = append(r, -1)
		}
		height = append(height, r)
	}
	queue := []coor1765{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				queue = append(queue, coor1765{i, j})
				height[i][j] = 0
			}
		}
	}

	for len(queue) > 0 {
		x, y := queue[0].x, queue[0].y
		queue = queue[1:]
		if x-1 >= 0 && height[x-1][y] == -1 {
			height[x-1][y] = height[x][y] + 1
			queue = append(queue, coor1765{x - 1, y})
		}
		if x+1 < m && height[x+1][y] == -1 {
			height[x+1][y] = height[x][y] + 1
			queue = append(queue, coor1765{x + 1, y})
		}
		if y-1 >= 0 && height[x][y-1] == -1 {
			height[x][y-1] = height[x][y] + 1
			queue = append(queue, coor1765{x, y - 1})
		}
		if y+1 < n && height[x][y+1] == -1 {
			height[x][y+1] = height[x][y] + 1
			queue = append(queue, coor1765{x, y + 1})
		}
	}

	return height
}
