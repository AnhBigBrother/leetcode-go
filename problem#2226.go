package main

func maximumCandies(candies []int, k int64) int {
	n := len(candies)
	l, r := 1, candies[0]
	for i := 1; i < n; i++ {
		if candies[i] > r {
			r = candies[i]
		}
	}
	for l < r {
		m := (l + r + 1) / 2
		sum := int64(0)
		for i := n - 1; i >= 0; i-- {
			sum += int64(candies[i] / m)
		}
		if sum >= k {
			l = m
		} else {
			r = m - 1
		}
	}
	sum := int64(0)
	for i := n - 1; i >= 0; i-- {
		sum += int64(candies[i] / r)
	}
	if sum < k {
		return 0
	}
	return r
}
