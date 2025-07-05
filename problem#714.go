package main

func maxProfit714(prices []int, fee int) int {
	buy, sell := 0, 1
	n := len(prices)

	dp := make([][2]int, n+1)
	dp[0][buy] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][buy] = max(dp[i-1][buy], dp[i-1][sell]-prices[i])
		dp[i][sell] = max(dp[i-1][sell], dp[i-1][buy]+prices[i]-fee)
	}

	return dp[n-1][sell]
}
