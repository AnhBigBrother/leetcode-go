package main

import "slices"

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	stack := []int{0}
	slices.SortFunc(envelopes, func(a, b []int) int {
		if a[0] > b[0] {
			return 1
		}
		if a[0] < b[0] {
			return -1
		}
		if a[1] > b[1] {
			return -1
		}
		return 1
	})
	for i := 1; i < n; i++ {
		if envelopes[stack[len(stack)-1]][0] < envelopes[i][0] && envelopes[stack[len(stack)-1]][1] < envelopes[i][1] {
			stack = append(stack, i)
		} else {
			l, r := 0, len(stack)-1
			for l < r {
				m := (l + r) / 2
				if envelopes[stack[m]][0] < envelopes[i][0] && envelopes[stack[m]][1] < envelopes[i][1] {
					l = m + 1
				} else {
					r = m
				}
			}
			stack[l] = i
		}
	}
	return len(stack)
}
