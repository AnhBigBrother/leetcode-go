package main

func maxScoreSightseeingPair(values []int) int {
	ans := 0
	maxSumValPos := values[0]
	for i := 1; i < len(values); i++ {
		ans = max(ans, maxSumValPos+values[i]-i)
		maxSumValPos = max(maxSumValPos, values[i]+i)
	}
	return ans
}
