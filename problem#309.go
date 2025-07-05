package main

func maxProfit309(prices []int) int {
	const buy, sell, cool int = 0, 1, 2
	n := len(prices)
	dp := [][]int{}
	for range n {
		dp = append(dp, make([]int, 3))
	}
	dp[0][buy] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][buy] = max(dp[i-1][buy], dp[i-1][cool]-prices[i])
		dp[i][sell] = max(dp[i-1][sell], dp[i-1][buy]+prices[i])
		dp[i][cool] = max(dp[i-1][cool], dp[i-1][sell])
	}
	return max(dp[n-1][cool], dp[n-1][sell])
}
