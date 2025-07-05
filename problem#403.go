package main

func canCross(stones []int) bool {
	n := len(stones)
	dp := []map[int]bool{}
	for range n {
		dp = append(dp, map[int]bool{})
	}
	dp[0][0] = true
	idx := map[int]int{}
	for i, p := range stones {
		idx[p] = i
	}
	for i, p := range stones {
		for k := range dp[i] {
			if i, ok := idx[p+k]; ok {
				dp[i][k] = true
			}
			if i, ok := idx[p+k+1]; ok {
				dp[i][k+1] = true
			}
			if i, ok := idx[p+k-1]; k-1 > 0 && ok {
				dp[i][k-1] = true
			}
		}
	}
	for _, b := range dp[n-1] {
		if b {
			return true
		}
	}
	return false
}
