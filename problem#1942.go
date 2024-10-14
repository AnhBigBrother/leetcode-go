package main

import (
	"container/heap"
	"slices"
)

type ChairNode struct {
	chair_number int
	time_leave   int
	index        int
}
type ChairHeap []*ChairNode

func (h ChairHeap) Len() int           { return len(h) }
func (h ChairHeap) Less(i, j int) bool { return h[i].time_leave < h[j].time_leave }
func (h ChairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *ChairHeap) Push(x any) {
	n := len(*h)
	item := x.(*ChairNode)
	item.index = n
	*h = append(*h, item)
}

func (h *ChairHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[0 : n-1]
	return item
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	fr := []int{}
	for i := 0; i < n; i++ {
		fr = append(fr, i)
	}
	slices.SortStableFunc(fr, func(i, j int) int {
		if times[i][0] > times[j][0] {
			return 1
		}
		return -1
	})
	q := ChairHeap{}
	un_occupied := MinHeap{}
	l := 0
	heap.Init(&q)
	heap.Init(&un_occupied)
	for _, x := range fr {
		for q.Len() > 0 && q[0].time_leave <= times[x][0] {
			p := heap.Pop(&q).(*ChairNode)
			heap.Push(&un_occupied, p.chair_number)
		}
		if un_occupied.Len() == 0 {
			if x == targetFriend {
				return l
			}
			heap.Push(&q, &ChairNode{
				chair_number: l,
				time_leave:   times[x][1],
			})
			l++
		} else {
			c := heap.Pop(&un_occupied).(int)
			if x == targetFriend {
				return c
			}
			heap.Push(&q, &ChairNode{
				chair_number: c,
				time_leave:   times[x][1],
			})
		}
	}
	return -1
}

// test: fmt.Println(smallestChair([][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}}, 6))
