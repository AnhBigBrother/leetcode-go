package main

import "slices"

func findLongestChain(pairs [][]int) int {
	slices.SortFunc(pairs, func(p1, p2 []int) int {
		if p1[0] == p2[0] {
			if p1[1] < p2[1] {
				return -1
			}
			return 1
		}
		if p1[0] < p2[0] {
			return -1
		}
		return 1
	})

	n := len(pairs)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	for i := 1; i < len(pairs); i++ {
		for j := range i {
			if pairs[j][1] < pairs[i][0] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	return dp[n-1]
}
