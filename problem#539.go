package main

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	minutes := []int{}
	for _, time := range timePoints {
		h, err1 := strconv.Atoi(time[:2])
		m, err2 := strconv.Atoi(time[3:5])
		if err1 == nil && err2 == nil {
			minutes = append(minutes, h*60+m)
		}
	}
	sort.Ints(minutes)
	n := len(minutes)
	ans := min(minutes[n-1]-minutes[0], minutes[0]+24*60-minutes[n-1])
	for i := 1; i < len(minutes); i++ {
		ans = min(ans, minutes[i]-minutes[i-1])
	}
	return ans
}
