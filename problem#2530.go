package main

import (
	"container/heap"
	"math"
)

func maxKelements(nums []int, k int) int64 {
	h := MaxHeap{}
	heap.Init(&h)
	for _, x := range nums {
		heap.Push(&h, x)
	}
	ans := 0
	for k > 0 {
		x := heap.Pop(&h).(int)
		ans += x
		heap.Push(&h, int(math.Ceil(float64(x)/3)))
		k--
	}
	return int64(ans)
}
