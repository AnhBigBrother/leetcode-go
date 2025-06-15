package main

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, 1001)
	}
	ans := 1
	for i := range n {
		for j := range i {
			diff := nums[i] - nums[j] + 500
			if dp[i][diff] < dp[j][diff]+1 {
				dp[i][diff] = dp[j][diff] + 1
				if dp[i][diff]+1 > ans {
					ans = dp[i][diff] + 1
				}
			}
		}
	}
	return ans
}
