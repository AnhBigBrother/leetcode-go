package main

func longestSubsequence(arr []int, difference int) int {
	n := len(arr)
	dp := make([]int, 20001)
	ans := 1
	for i := range n {
		x := arr[i] + 10000
		y := x - difference
		if y >= 0 && y < len(dp) {
			if dp[x] < dp[y]+1 {
				dp[x] = dp[y] + 1
			}
		} else {
			dp[x] = 1
		}
		if ans < dp[x] {
			ans = dp[x]
		}
	}
	return ans
}
