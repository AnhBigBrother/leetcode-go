package main

func maxProfit123(prices []int) int {
	buy1, buy2, sell1, sell2 := -prices[0], -prices[0], 0, 0
	n := len(prices)
	for i := 1; i < n; i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return max(sell1, sell2)
}
