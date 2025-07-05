package main

import "slices"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		dp[a] = -1
	}
	valid_coins := []int{}
	for _, c := range coins {
		if c <= amount {
			valid_coins = append(valid_coins, c)
		}
	}
	slices.Sort(valid_coins)
	for _, c := range valid_coins {
		dp[c] = 1
	}
	for a := 1; a <= amount; a++ {
		for i := range valid_coins {
			c := valid_coins[i]
			if a-c < 0 {
				break
			}
			if dp[a-c] == -1 {
				continue
			}
			if dp[a] == -1 {
				dp[a] = dp[a-c] + 1
			} else {
				dp[a] = min(dp[a], dp[a-c]+1)
			}
		}
	}
	return dp[amount]
}
