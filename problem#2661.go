package main

type coor2661 struct {
	x int
	y int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	sumRow := make([]int, m)
	sumCol := make([]int, n)
	pos := make([]coor2661, m*n+1)

	for i, r := range mat {
		for j, c := range r {
			sumRow[i] += c
			sumCol[j] += c
			pos[c] = coor2661{i, j}
		}
	}

	for i, num := range arr {
		x, y := pos[num].x, pos[num].y
		sumRow[x] -= num
		sumCol[y] -= num
		if sumRow[x] == 0 || sumCol[y] == 0 {
			return i
		}
	}

	return arr[len(arr)-1]
}
