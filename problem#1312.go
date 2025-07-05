package main

func minInsertions(s string) int {
	n := len(s)
	dp := [][]int{}
	for range n + 1 {
		dp = append(dp, make([]int, n+1))
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == s[n-j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return n - dp[n][n]
}
