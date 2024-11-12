package main

import (
	"math"
	"slices"
)

func maximumBeauty(items [][]int, queries []int) []int {
	beauty := map[int]int{}
	price := []int{}
	for _, x := range items {
		if beauty[x[0]] == 0 {
			price = append(price, x[0])
			beauty[x[0]] = x[1]
		} else if x[1] > beauty[x[0]] {
			beauty[x[0]] = x[1]
		}
	}
	slices.Sort(price)
	max_beauty := map[int]int{}
	max_beauty[price[0]] = beauty[price[0]]
	for i := 1; i < len(price); i++ {
		max_beauty[price[i]] = max(max_beauty[price[i-1]], beauty[price[i]])
	}
	binSearch := func(x int) int {
		l, r := 0, len(price)-1
		for l < r {
			m := int(math.Ceil(float64(l+r) / 2))
			if price[m] > x {
				r = m - 1
			} else {
				l = m
			}
		}
		if price[r] <= x {
			return max_beauty[price[r]]
		}
		return 0
	}
	ans := []int{}
	for _, x := range queries {
		ans = append(ans, binSearch(x))
	}
	return ans
}
