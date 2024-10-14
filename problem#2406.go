package main

import (
	"container/heap"
	"slices"
)

func minGroups(intervals [][]int) int {
	slices.SortStableFunc(intervals, func(i, j []int) int {
		if i[0] < j[0] {
			return -1
		} else if i[0] > j[0] {
			return 1
		} else {
			if i[1] < j[1] {
				return -1
			}
			return 1
		}
	})
	// fmt.Println(intervals)
	ans := 1
	mh := MinHeap{intervals[0][1]}
	heap.Init(&mh)
	for i := 1; i < len(intervals); i++ {
		x := intervals[i]
		if mh[0] >= x[0] {
			heap.Push(&mh, x[1])
			ans++
		} else {
			heap.Pop(&mh)
			heap.Push(&mh, x[1])
		}
	}
	return ans
}
