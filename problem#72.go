package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := [][]int{} // dp[i][j] = minDistance(word1[:i], word2[:j])
	for range m + 1 {
		dp = append(dp, make([]int, n+1))
	}
	for i := range m + 1 {
		dp[i][0] = i
	}
	for j := range n + 1 {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 1
			if word1[i-1] == word2[j-1] {
				cost = 0
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}
	return dp[m][n]
}
