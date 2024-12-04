package main

func getLargestOutlier(nums []int) int {
	sum := 0
	count := map[int]int{}
	for _, x := range nums {
		sum += x
		count[x]++
	}
	ans := -1001
	for _, x := range nums {
		ssum := sum - x
		count[x]--
		if ssum%2 == 0 {
			ssum /= 2
			if count[ssum] > 0 {
				ans = max(ans, x)
			}
		}
		count[x]++
	}
	return ans
}
