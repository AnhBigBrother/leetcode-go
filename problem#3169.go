package main

import (
	"slices"
)

func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		if a[0] == b[0] {
			if a[1] < b[1] {
				return -1
			}
			return 1
		}
		if a[0] < b[0] {
			return -1
		}
		return 1
	})
	windows := [][]int{}
	window := []int{meetings[0][0], meetings[0][1]}
	for i := 1; i < len(meetings); i++ {
		w := meetings[i]
		if w[0] <= window[1] {
			if window[1] < w[1] {
				window[1] = w[1]
			}
		} else {
			temp := []int{window[0], window[1]}
			windows = append(windows, temp)
			window[0] = w[0]
			window[1] = w[1]
		}
	}
	windows = append(windows, window)
	ans := days - windows[len(windows)-1][1]
	cur := 0

	for _, x := range windows {
		ans += x[0] - cur - 1
		cur = x[1]
	}

	return ans
}
