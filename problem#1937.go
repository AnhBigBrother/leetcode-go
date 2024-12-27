package main

func maxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := []int{}
	dp = append(dp, points[0]...)
	for i := 1; i < m; i++ {
		row := points[i]
		left, right := make([]int, n), make([]int, n)
		left[0], right[n-1] = dp[0], dp[n-1]-n+1
		for l := 1; l < n; l++ {
			left[l] = max(left[l-1], dp[l]+l)
		}
		for r := n - 2; r >= 0; r-- {
			right[r] = max(right[r+1], dp[r]-r)
		}
		for j := 0; j < n; j++ {
			dp[j] = max(row[j]-j+left[j], row[j]+j+right[j])
		}
	}
	ans := dp[0]
	for i := 1; i < n; i++ {
		ans = max(ans, dp[i])
	}
	return int64(ans)
}
