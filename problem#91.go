package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	if n == 1 {
		return 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' && s[i-1] == '0' {
			return 0
		}
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 0
	if s[1] != '0' {
		dp[1] += 1
	}
	x := (rune(s[0])-48)*10 + rune(s[1]) - 48
	if x <= 26 {
		dp[1] += 1
	}
	for i := 2; i < n; i++ {
		if s[i] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-1] != '0' {
			x := (rune(s[i-1])-48)*10 + rune(s[i]) - 48
			if x <= 26 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n-1]
}
