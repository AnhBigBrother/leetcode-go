package main

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	var findLeng func(x, y int) int
	findLeng = func(x, y int) int {
		if dp[x][y] > 0 {
			return dp[x][y]
		}
		l, r, u, d := 0, 0, 0, 0
		if x-1 >= 0 && matrix[x-1][y] > matrix[x][y] {
			u = findLeng(x-1, y)
		}
		if x+1 < m && matrix[x+1][y] > matrix[x][y] {
			d = findLeng(x+1, y)
		}
		if y-1 >= 0 && matrix[x][y-1] > matrix[x][y] {
			l = findLeng(x, y-1)
		}
		if y+1 < n && matrix[x][y+1] > matrix[x][y] {
			r = findLeng(x, y+1)
		}
		dp[x][y] = max(l, r, u, d, 0) + 1
		return dp[x][y]
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, findLeng(i, j))
		}
	}
	return ans
}
