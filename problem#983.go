package main

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	isDay := map[int]bool{}
	for _, d := range days {
		isDay[d] = true
	}
	dp := make([]int, days[n-1]+1)
	for i := 1; i <= days[n-1]; i++ {
		if !isDay[i] {
			dp[i] = dp[i-1]
		} else {
			x := i - 1
			y := i - 7
			z := i - 30
			if y < 0 {
				y = 0
			}
			if z < 0 {
				z = 0
			}
			dp[i] = min(dp[x]+costs[0], dp[y]+costs[1], dp[z]+costs[2])
		}
	}
	return dp[days[n-1]]
}
