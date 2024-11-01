package main

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	ans := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = matrix[i][0]
		ans += dp[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = matrix[0][j]
		ans += dp[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				ans += dp[i][j]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ans
}
