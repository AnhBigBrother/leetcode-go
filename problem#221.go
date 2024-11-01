package main

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	ans := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = 1
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ans = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				ans = max(ans, dp[i][j]*dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}

	return ans
}
