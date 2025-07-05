package main

func findMaxForm(strs []string, m int, n int) int {
	dp := [][]int{}
	for range m + 1 {
		dp = append(dp, make([]int, n+1))
	}
	countZeroOne := func(s string) (int, int) {
		z := 0
		for _, c := range s {
			if c == '0' {
				z++
			}
		}
		return z, len(s) - z
	}
	for _, s := range strs {
		z, o := countZeroOne(s)
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = max(dp[i][j], dp[i-z][j-o]+1)
			}
		}
	}
	return dp[m][n]
}
