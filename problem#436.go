package main

import "slices"

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	ans := make([]int, n)
	start := make([]int, n)
	idx_map := map[int]int{}
	for i, x := range intervals {
		start[i] = x[0]
		idx_map[x[0]] = i
	}
	slices.Sort(start)
	for i := range n {
		target := intervals[i][1]
		l, r := 0, n-1
		for l < r {
			m := (l + r) / 2
			if start[m] < target {
				l = m + 1
			} else {
				r = m
			}
		}
		if start[l] < target {
			ans[i] = -1
		} else {
			ans[i] = idx_map[start[l]]
		}
	}
	return ans
}
