package main

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	ans := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		l := 0
		if grid[i][0] < grid[i][1] {
			l = dp[i][0] + 1
		}
		if i-1 >= 0 && grid[i-1][0] < grid[i][1] {
			l = max(l, dp[i-1][0]+1)
		}
		if i+1 < m && grid[i+1][0] < grid[i][1] {
			l = max(l, dp[i+1][0]+1)
		}
		if l != 0 {
			dp[i][1] = l
			ans = max(ans, dp[i][1])
		}
	}
	for j := 2; j < n; j++ {
		for i := 0; i < m; i++ {
			l := 0
			if dp[i][j-1] != 0 && grid[i][j-1] < grid[i][j] {
				l = dp[i][j-1] + 1
			}
			if i-1 >= 0 && dp[i-1][j-1] != 0 && grid[i-1][j-1] < grid[i][j] {
				l = max(l, dp[i-1][j-1]+1)
			}
			if i+1 < m && dp[i+1][j-1] != 0 && grid[i+1][j-1] < grid[i][j] {
				l = max(l, dp[i+1][j-1]+1)
			}
			if l != 0 {
				dp[i][j] = l
				ans = max(ans, l)
			}
		}
	}
	return ans
}
