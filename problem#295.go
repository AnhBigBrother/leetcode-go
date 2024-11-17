package main

import "container/heap"

type MedianFinder struct {
	h1 MaxHeap
	h2 MinHeap
}

func Constructor295() MedianFinder {
	h1 := MaxHeap{}
	h2 := MinHeap{}
	heap.Init(&h1)
	heap.Init(&h2)
	return MedianFinder{
		h1: h1,
		h2: h2,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.h2, num)
	if this.h1.Len() == this.h2.Len() {
		if this.h1[0] > this.h2[0] {
			heap.Push(&this.h1, heap.Pop(&this.h2))
			heap.Push(&this.h2, heap.Pop(&this.h1))
		}
	} else {
		heap.Push(&this.h1, heap.Pop(&this.h2))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	l, r := this.h1.Len(), this.h2.Len()
	if l > r {
		return float64(this.h1[0])
	}
	return float64(this.h1[0]+this.h2[0]) / 2
}
