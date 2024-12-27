package main

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < int(math.Abs(float64(target))) {
		return 0
	}
	sum += target
	if sum%2 == 1 {
		return 0
	}
	sum /= 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for x := sum; x >= num; x-- {
			dp[x] += dp[x-num]
		}
	}
	return dp[sum]
}
