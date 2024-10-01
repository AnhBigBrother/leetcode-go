package main

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	dp := []int{}
	dp = append(dp, 0)
	for i := 1; i <= n; i++ {
		dp = append(dp, 100)
	}
	inDict := make(map[string]bool)
	for _, x := range dictionary {
		inDict[x] = true
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if inDict[s[j:i]] {
				dp[i] = min(dp[i], dp[j])
			} else {
				dp[i] = min(dp[i], dp[j]+len(s[j:i]))
			}
		}
	}
	return dp[n]
}
