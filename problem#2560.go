package main

import "slices"

func minCapability(nums []int, k int) int {
	n := len(nums)
	vals := []int{}
	vals = append(vals, nums...)
	slices.Sort(vals)
	l, r := 0, n-1
	for l < r {
		m := (l + r) / 2
		count := 0
		f := false
		for _, x := range nums {
			if f {
				f = false
				continue
			}
			if x <= vals[m] {
				f = true
				count++
			}
		}
		if count >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return vals[l]
}
