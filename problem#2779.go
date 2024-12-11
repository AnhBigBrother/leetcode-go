package main

import (
	"math"
	"slices"
)

func maximumBeauty(nums []int, k int) int {
	n := len(nums)
	slices.Sort(nums)
	minNum := nums[0]
	maxNum := nums[len(nums)-1]
	findFirst := func(num int) int {
		l, r := 0, n-1
		for l < r {
			m := (l + r) / 2
			if nums[m]+k < num {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}
	findLast := func(num int) int {
		l, r := 0, n-1
		for l < r {
			m := int(math.Ceil(float64(l+r) / 2))
			if nums[m]-k > num {
				r = m - 1
			} else {
				l = m
			}
		}
		return r
	}
	ans := 0
	for x := minNum; x <= maxNum; x++ {
		l := findFirst(x)
		r := findLast(x)
		ans = max(ans, r-l+1)
	}
	return ans
}
