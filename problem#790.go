package main

func numTilings(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 5
	}
	mod := int(1e9 + 7)
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1]*2 + dp[i-3]
		dp[i] %= mod
	}
	return dp[n]
}
