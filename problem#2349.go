package main

import "container/heap"

// priority queue definition
type Item2349 struct {
	Value      int
	Key        int
	Heap_Index int
}

type PQ2349 []*Item2349

func (pq PQ2349) Len() int { return len(pq) }

func (pq PQ2349) Less(i, j int) bool {
	if pq[i].Key == -1 && pq[j].Key == -1 {
		return false
	}
	if pq[i].Key == -1 {
		return false
	}
	if pq[j].Key == -1 {
		return true
	}
	return pq[i].Key < pq[j].Key
}

func (pq PQ2349) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Heap_Index = i
	pq[j].Heap_Index = j
}

func (pq *PQ2349) Push(x any) {
	n := len(*pq)
	item := x.(*Item2349)
	item.Heap_Index = n
	*pq = append(*pq, item)
}

func (pq *PQ2349) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Heap_Index = -1
	*pq = old[0 : n-1]
	return item
}

// NumberContainers data structure
type NumberContainers struct {
	Key_Map  map[int]*PQ2349
	Item_Map map[int]*Item2349
}

func NewNumberContainers() NumberContainers {
	return NumberContainers{
		Key_Map:  map[int]*PQ2349{},
		Item_Map: map[int]*Item2349{},
	}
}

func (this *NumberContainers) Change(key int, val int) {
	old_item := this.Item_Map[key]
	if old_item != nil {
		old_item.Key = -1
		heap.Fix(this.Key_Map[old_item.Value], old_item.Heap_Index)
	}
	if this.Key_Map[val] == nil {
		this.Key_Map[val] = &PQ2349{}
		heap.Init(this.Key_Map[val])
	}
	new_item := &Item2349{
		Value: val,
		Key:   key,
	}
	heap.Push(this.Key_Map[val], new_item)
	this.Item_Map[key] = new_item
}

func (this *NumberContainers) Find(val int) int {
	km := this.Key_Map[val]
	if km == nil || km.Len() == 0 {
		return -1
	}
	return (*km)[0].Key
}
