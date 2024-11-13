package main

import "slices"

func deleteAndEarn(nums []int) int {
	amount := map[int]int{}
	distinct_nums := []int{}
	for _, x := range nums {
		if amount[x] == 0 {
			distinct_nums = append(distinct_nums, x)
		}
		amount[x] += x
	}
	slices.Sort(distinct_nums)
	n := len(distinct_nums)
	if n == 1 {
		return amount[distinct_nums[0]]
	}
	dp := make([]int, n)
	dp[0] = amount[distinct_nums[0]]
	if distinct_nums[1] == distinct_nums[0]+1 {
		dp[1] = max(amount[distinct_nums[0]], amount[distinct_nums[1]])
	} else {
		dp[1] = dp[0] + amount[distinct_nums[1]]
	}
	for i := 2; i < n; i++ {
		if distinct_nums[i] == distinct_nums[i-1]+1 {
			dp[i] = max(dp[i-2]+amount[distinct_nums[i]], dp[i-1])
		} else {
			dp[i] = dp[i-1] + amount[distinct_nums[i]]
		}
	}
	return dp[n-1]
}
