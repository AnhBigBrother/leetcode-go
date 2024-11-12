package main

import "slices"

func largestCombination(candidates []int) int {
	maxC := slices.Max(candidates)
	ans := 0
	for x := 1; x <= maxC; x <<= 1 {
		count := 0
		for _, c := range candidates {
			if c&x > 0 {
				count++
			}
		}
		ans = max(ans, count)
	}
	return ans
}
