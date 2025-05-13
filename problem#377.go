package main

import "slices"

func combinationSum4(nums []int, target int) int {
	slices.Sort(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for x := 1; x <= target; x++ {
		for i := range nums {
			if nums[i] > x {
				break
			}
			dp[x] += dp[x-nums[i]]
		}
	}
	return dp[target]
}
