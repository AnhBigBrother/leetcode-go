package main

import "container/heap"

type LFUItem struct {
	Key  int
	Val  int
	Freq int
	Last int
	Idx  int
}

type LFUQueue []*LFUItem

func (pq LFUQueue) Len() int { return len(pq) }

func (pq LFUQueue) Less(i, j int) bool {
	if pq[i].Freq == pq[j].Freq {
		return pq[i].Last < pq[j].Last
	}
	return pq[i].Freq < pq[j].Freq
}

func (pq LFUQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Idx = i
	pq[j].Idx = j
}

func (pq *LFUQueue) Push(x any) {
	n := len(*pq)
	item := x.(*LFUItem)
	item.Idx = n
	*pq = append(*pq, item)
}

func (pq *LFUQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Idx = -1
	*pq = old[0 : n-1]
	return item
}

type LFUCache struct {
	Cache LFUQueue
	LMap  map[int]*LFUItem
	Cap   int
	Last  int
}

func LFUConstructor(capacity int) LFUCache {
	x := LFUCache{
		Cache: LFUQueue{},
		LMap:  map[int]*LFUItem{},
		Cap:   capacity,
		Last:  0,
	}
	heap.Init(&(x.Cache))
	return x
}

func (this *LFUCache) Get(key int) int {
	x := this.LMap[key]
	if x == nil {
		return -1
	}
	this.Last++
	x.Last = this.Last
	x.Freq++
	heap.Fix(&(this.Cache), x.Idx)
	return x.Val
}

func (this *LFUCache) Put(key int, value int) {
	if this.LMap[key] == nil {
		if this.Cache.Len() == this.Cap {
			x := heap.Pop(&(this.Cache)).(*LFUItem)
			this.LMap[x.Key] = nil
		}
		this.Last++
		newItem := LFUItem{
			Key:  key,
			Val:  value,
			Freq: 1,
			Last: this.Last,
			Idx:  len(this.Cache),
		}
		this.LMap[key] = &newItem
		heap.Push(&(this.Cache), &newItem)
	} else {
		x := this.LMap[key]
		x.Freq++
		this.Last++
		x.Last = this.Last
		x.Val = value
		heap.Fix(&(this.Cache), x.Idx)
	}
}
