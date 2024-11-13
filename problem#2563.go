package main

import (
	"math"
	"slices"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	slices.Sort(nums)
	binSearch := func(i int) int {
		l, r := i+1, n-1
		for l < r {
			m := int(math.Floor(float64(l+r) / 2))
			if nums[m]+nums[i] >= lower {
				r = m
			} else {
				l = m + 1
			}
		}
		if l > r || nums[i]+nums[l] < lower || nums[i]+nums[l] > upper {
			return 0
		}
		left := l
		l, r = i+1, n-1
		for l < r {
			m := int(math.Ceil(float64(l+r) / 2))
			if nums[m]+nums[i] <= upper {
				l = m
			} else {
				r = m - 1
			}
		}
		if l > r || nums[i]+nums[r] < lower || nums[i]+nums[r] > upper {
			return 0
		}
		right := r
		return right - left + 1
	}
	ans := int64(0)
	for i := 0; i < n; i++ {
		ans += int64(binSearch(i))
	}
	return ans
}
