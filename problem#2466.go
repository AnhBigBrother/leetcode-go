package main

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high)
	ans := 0
	mod := int(1e9 + 7)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
		if i >= low {
			ans = (ans + dp[i]) % mod
		}
	}
	return ans
}
