package main

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := [][]int{}
	for range n {
		dp = append(dp, make([]int, n))
	}
	for i := range n {
		dp[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j+1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j+1])
			}
		}
	}
	return dp[n-1][0]
}
