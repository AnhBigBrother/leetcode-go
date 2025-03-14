package main

import "slices"

func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	ans := []byte{}
	i, j := m, n
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			ans = append(ans, str1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			ans = append(ans, str1[i-1])
			i--
		} else {
			ans = append(ans, str2[j-1])
			j--
		}
	}
	for i > 0 {
		ans = append(ans, str1[i-1])
		i--
	}
	for j > 0 {
		ans = append(ans, str2[j-1])
		j--
	}
	slices.Reverse(ans)
	return string(ans)
}
