package main

import "slices"

func maxTwoEvents(events [][]int) int {
	ans := 0
	mx := map[int]int{} // mx[a] = max value of all duration i which start[i] == a
	for _, e := range events {
		mx[e[0]] = max(mx[e[0]], e[2])
	}
	starts := []int{}
	for key := range mx {
		starts = append(starts, key)
	}
	slices.Sort(starts)
	n := len(starts)
	mRight := map[int]int{}
	m := 0
	for i := n - 1; i >= 0; i-- {
		m = max(m, mx[starts[i]])
		mRight[starts[i]] = m
	}
	find := func(a int) int {
		l, r := 0, n-1
		for l < r {
			if a > starts[r] {
				return -1
			}
			m := (l + r) / 2
			if starts[m] >= a {
				r = m
			} else {
				l = m + 1
			}
		}
		if a > starts[r] {
			return -1
		}
		return starts[r]
	}
	for _, e := range events {
		r := find(e[1] + 1)
		sum := e[2] + mRight[r]
		ans = max(ans, sum)
	}
	return ans
}
