package main

import (
	"container/heap"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	h := MaxHeap{}
	for _, g := range gifts {
		h = append(h, g)
	}
	heap.Init(&h)
	for k > 0 {
		x := heap.Pop(&h).(int)
		x = int(math.Floor(math.Sqrt(float64(x))))
		heap.Push(&h, x)
		k--
	}
	ans := 0
	for _, x := range h {
		ans += x
	}
	return int64(ans)
}
