package main

import "math"

func minimumSize(nums []int, maxOperations int) int {
	sum := 0
	l, r := 1, nums[0]
	for _, x := range nums {
		sum += x
		r = max(r, x)
	}
	maxLen := len(nums) + maxOperations
	for l < r {
		m := (l + r) / 2
		ops := 0
		for _, num := range nums {
			ops += int(math.Ceil(float64(num) / float64(m)))
		}
		if ops <= maxLen {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}
