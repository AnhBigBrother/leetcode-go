package main

import "container/heap"

func minOperations3066(nums []int, k int) int {
	h := MinHeap{}
	h = append(h, nums...)
	heap.Init(&h)
	ans := 0
	for h.Len() >= 2 && h[0] < k {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		a := min(x, y)*2 + max(x, y)
		heap.Push(&h, a)
		ans++
	}
	return ans
}
