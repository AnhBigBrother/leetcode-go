package main

import "container/heap"

type item3264 struct {
	Value int
	Idx   int
	index int
}

type pQueue3264 []*item3264

func (pq pQueue3264) Len() int { return len(pq) }

func (pq pQueue3264) Less(i, j int) bool {
	if pq[i].Value == pq[j].Value {
		return pq[i].Idx < pq[j].Idx
	}
	return pq[i].Value < pq[j].Value
}

func (pq pQueue3264) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pQueue3264) Push(x any) {
	n := len(*pq)
	item := x.(*item3264)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pQueue3264) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func getFinalState(nums []int, k int, multiplier int) []int {
	pq := pQueue3264{}
	for i, x := range nums {
		pq = append(pq, &item3264{Value: x, Idx: i})
	}
	heap.Init(&pq)
	for k > 0 {
		x := heap.Pop(&pq).(*item3264)
		x.Value *= multiplier
		heap.Push(&pq, x)
		k--
	}
	ans := make([]int, len(nums))
	for _, item := range pq {
		ans[item.Idx] = item.Value
	}
	return ans
}
