package main

func maxProfit188(k int, prices []int) int {
	n := len(prices)
	buy := make([]int, k)
	sell := make([]int, k)
	for i := range k {
		buy[i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], -prices[i])
		for j := 1; j < k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
		}
		for j := range k {
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k-1]
}
